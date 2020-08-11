package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const cid2InfoUri = "/api/thirdPart/getCid2Info"

// Cid2Info
type Cid2Info struct {
	BaseClient
	Token string
}

// Cid2InfoResponse
type Cid2InfoResponse struct {
	ErrorResponse
	Data []struct {
		CID2   int    `json:"cid2"`
		CName2 string `json:"cname2"`
		PicUrl string `json:"pic_url"`
	} `json:"data;omitempty"`
}

// Handle
func (p *Cid2Info) Handle(timestamp string) (*Cid2InfoResponse, error) {
	return p.do(DouYuDomain+cid2InfoUri, timestamp)
}

// do
func (p *Cid2Info) do(url, timestamp string) (*Cid2InfoResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, cid2InfoUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, ""); err != nil {
		return nil, err
	} else {
		var ret, errResp = new(Cid2InfoResponse), new(ErrorResponse)
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
