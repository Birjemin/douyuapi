package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const liveUri = "/api/thirdPart/Live"

// Live
type Live struct {
	BaseClient
	Token string
}

// LiveResponse
type LiveResponse struct {
	ErrorResponse
	Data struct {
		RID        int         `json:"rid"`
		RoomName   string      `json:"room_name"`
		LiveUrl    string      `json:"live_url"`
		HlsUrl     string      `json:"hls_url"`
		MixUrl     string      `json:"mix_url"`
		RateSwitch int         `json:"rate_switch"`
		ShowStatus int         `json:"show_status"`
		HlsMul     interface{} `json:"hls_mul"`
		FlvMul     interface{} `json:"flv_mul"`
	} `json:"data"`
}

// Handle
func (p *Live) Handle(postJson, timestamp string) (*LiveResponse, error) {
	return p.do(DouYuDomain+liveUri, postJson, timestamp)
}

// do
func (p *Live) do(url string, postJson string, timestamp string) (*LiveResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, liveUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJson); err != nil {
		return nil, err
	} else {
		var ret = new(LiveResponse)
		if err := p.Client.GetResponseJson(ret); err != nil {
			return nil, err
		} else {
			return ret, nil
		}

	}
}
