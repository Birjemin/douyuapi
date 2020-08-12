package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const getPlayURI = "/api/thirdPart/getPlay"

// GetPlay ...
type GetPlay struct {
	BaseClient
	Token string
}

// GetPlayResponse ...
type GetPlayResponse struct {
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
func (p *GetPlay) Handle(postJSON, timestamp string) (*GetPlayResponse, error) {
	return p.do(DouYuDomain+getPlayURI, postJSON, timestamp)
}

// do
func (p *GetPlay) do(url, postJSON, timestamp string) (*GetPlayResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, getPlayURI, params)

	url += "?" + utils.HTTPQueryBuild(params)

	if err := p.Client.HTTPPostJSON(url, postJSON); err != nil {
		return nil, err
	}
	var ret, errResp = new(GetPlayResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJSON(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
