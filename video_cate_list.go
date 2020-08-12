package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const videoCateListURI = "/api/thirdPart/video/cateList"

// VideoCateList ...
type VideoCateList struct {
	BaseClient
	Token string
}

// VideoCateListResponse ...
type VideoCateListResponse struct {
	ErrorResponse
	Data []struct {
		Cid2   int    `json:"cid2"`
		Cname2 string `json:"cname2"`
	} `json:"data"`
}

// Handle ...
func (p *VideoCateList) Handle(timestamp string) (*VideoCateListResponse, error) {
	return p.do(DouYuDomain+videoCateListURI, timestamp)
}

// do
func (p *VideoCateList) do(url, timestamp string) (*VideoCateListResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, videoCateListURI, params)

	url += "?" + utils.HTTPQueryBuild(params)

	if err := p.Client.HTTPPostJSON(url, "{}"); err != nil {
		return nil, err
	}
	var ret, errResp = new(VideoCateListResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJSON(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
