package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const vodDownUrlUri = "/api/thirdPart/video/vodDownUrl"

// VodDownUrl
type VodDownUrl struct {
	BaseClient
	Token string
}

// VodDownUrlResponse
type VodDownUrlResponse struct {
	ErrorResponse
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
}

// Handle
func (p *VodDownUrl) Handle(postJson, timestamp string) (*VodDownUrlResponse, error) {
	return p.do(DouYuDomain+vodDownUrlUri, postJson, timestamp)
}

// do
func (p *VodDownUrl) do(url, postJson, timestamp string) (*VodDownUrlResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, vodDownUrlUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJson); err != nil {
		return nil, err
	} else {
		var ret = new(VodDownUrlResponse)
		if err := p.Client.GetResponseJson(ret); err != nil {
			return nil, err
		} else {
			return ret, nil
		}
	}
}