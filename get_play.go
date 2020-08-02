package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const PlayUri = "/api/thirdPart/getPlay"

// Play
type Play struct {
	BaseClient
	Token string
}

// PlayResponse
type PlayResponse struct {
	ErrorResponse
	Data struct {
		RID        int         `json:"rid"`
		RoomName   string      `json:"room_name"`
		LiveUrl    string      `json:"live_url"`
		HlsUrl     string      `json:"hls_url"`
		MixUrl     string      `json:"mix_url"`
		RateSwitch int      `json:"rate_switch"`
		ShowStatus int      `json:"show_status"`
		HlsMul     interface{} `json:"hls_mul"`
		FlvMul     interface{} `json:"flv_mul"`
	} `json:"data"`
}

// Handle
func (p *Play) Handle(postJson, timestamp string) (*PlayResponse, error) {
	return p.do(DouYuDomain+PlayUri, postJson, timestamp)
}

// do
func (p *Play) do(url string, postJson string, timestamp string) (*PlayResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, PlayUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJson); err != nil {
		return nil, err
	} else {
		var ret = new(PlayResponse)
		if err = p.Client.GetResponseJson(ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
}
