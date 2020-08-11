package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const cid1InfoUri = "/api/thirdPart/getCid1Info"

// Cid1Info
type Cid1Info struct {
	BaseClient
	Token string
}

// Cid1InfoResponse
type Cid1InfoResponse struct {
	ErrorResponse
	Data []struct {
		CID1   int    `json:"cid1"`
		CName1 string `json:"cname1"`
	} `json:"data;omitempty"`
}

// Handle
func (p *Cid1Info) Handle(timestamp string) (*Cid1InfoResponse, error) {
	return p.do(DouYuDomain+cid1InfoUri, timestamp)
}

// do
func (p *Cid1Info) do(url, timestamp string) (*Cid1InfoResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, cid1InfoUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, ""); err != nil {
		return nil, err
	} else {
		var ret, errResp = new(Cid1InfoResponse), new(ErrorResponse)
		if err = p.Client.GetResponseJson(ret, errResp); err != nil {
			return nil, err
		}
		if errResp.Code != 0 {
			ret.Code = errResp.Code
			ret.Msg = errResp.Msg
		}
		return ret, nil
	}
}
