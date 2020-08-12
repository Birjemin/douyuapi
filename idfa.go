package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const idfaURI = "/api/thirdPart/idfa"

// IDFAInfo ...
type IDFAInfo struct {
	BaseClient
	Token string
}

// IDFAResponse ...
type IDFAResponse struct {
	ErrorResponse
	Data map[string]int `json:"data"`
}

// Handle ...
func (p *IDFAInfo) Handle(idfa, timestamp string) (*IDFAResponse, error) {
	return p.do(DouYuDomain+idfaURI, idfa, timestamp)
}

// do
func (p *IDFAInfo) do(url, idfa, timestamp string) (*IDFAResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, idfaURI, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, idfa); err != nil {
		return nil, err
	}
	var ret, errResp = new(IDFAResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJson(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
