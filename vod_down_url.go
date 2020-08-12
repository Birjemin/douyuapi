package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const vodDownURLURI = "/api/thirdPart/video/vodDownURL"

// VodDownURL ...
type VodDownURL struct {
	BaseClient
	Token string
}

// VodDownURLResponse ...
type VodDownURLResponse struct {
	ErrorResponse
	Data struct {
		URL string `json:"url"`
	} `json:"data"`
}

// Handle ...
func (p *VodDownURL) Handle(postJSON, timestamp string) (*VodDownURLResponse, error) {
	return p.do(DouYuDomain+vodDownURLURI, postJSON, timestamp)
}

// do
func (p *VodDownURL) do(url, postJSON, timestamp string) (*VodDownURLResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, vodDownURLURI, params)

	url += "?" + utils.HTTPQueryBuild(params)

	if err := p.Client.HTTPPostJSON(url, postJSON); err != nil {
		return nil, err
	}
	var ret, errResp = new(VodDownURLResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJSON(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
