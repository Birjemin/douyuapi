package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const playURI = "/api/thirdPart/getPlay"

// Play ...
type Play struct {
	BaseClient
	Token string
}

// PlayResponse ...
type PlayResponse struct {
	ErrorResponse
	Data struct {
		RID        int         `json:"rid"`
		RoomName   string      `json:"room_name"`
		LiveURL    string      `json:"live_url"`
		HlsURL     string      `json:"hls_url"`
		MixURL     string      `json:"mix_url"`
		RateSwitch int         `json:"rate_switch"`
		ShowStatus int         `json:"show_status"`
		HlsMul     interface{} `json:"hls_mul"`
		FlvMul     interface{} `json:"flv_mul"`
	} `json:"data"`
}

// Handle ...
func (p *Play) Handle(postJSON, timestamp string) (*PlayResponse, error) {
	return p.do(DouYuDomain+playURI, postJSON, timestamp)
}

// do
func (p *Play) do(url, postJSON, timestamp string) (*PlayResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, playURI, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJSON); err != nil {
		return nil, err
	}
	var ret, errResp = new(PlayResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJson(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
