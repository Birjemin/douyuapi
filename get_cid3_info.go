package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const cid3InfoURI = "/api/thirdPart/getCid3Info"

// Cid3Info ...
type Cid3Info struct {
	BaseClient
	Token string
}

// Cid3InfoResponse ...
type Cid3InfoResponse struct {
	ErrorResponse
	Data []struct {
		CID3   int    `json:"cid3"`
		CName3 string `json:"cname3"`
	} `json:"data"`
}

// Handle ...
func (p *Cid3Info) Handle(timestamp string) (*Cid3InfoResponse, error) {
	return p.do(DouYuDomain+cid3InfoURI, timestamp)
}

// do
func (p *Cid3Info) do(url, timestamp string) (*Cid3InfoResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, cid3InfoURI, params)

	url += "?" + utils.HTTPQueryBuild(params)

	if err := p.Client.HTTPPostJSON(url, ""); err != nil {
		return nil, err
	}
	var ret, errResp = new(Cid3InfoResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJSON(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
