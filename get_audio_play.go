package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const getAudioPlayUri = "/api/thirdPart/getAudioPlay"

// GetAudioPlay
type GetAudioPlay struct {
	BaseClient
	Token string
}

// GetAudioPlayResponse
type GetAudioPlayResponse struct {
	ErrorResponse
	Data []struct {
		RID      int    `json:"rid"`
		UID      int    `json:"uid"`
		RtmpCdn  string `json:"rtmpCdn"`
		RtmpLive string `json:"rtmpLive"`
		RtmpURL  string `json:"rtmpURL"`
	} `json:"data;omitempty"`
}

// Handle
func (p *GetAudioPlay) Handle(postJson, timestamp string) (*GetAudioPlayResponse, error) {
	return p.do(DouYuDomain+getAudioPlayUri, postJson, timestamp)
}

// do
func (p *GetAudioPlay) do(url, postJson, timestamp string) (*GetAudioPlayResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, getAudioPlayUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJson); err != nil {
		return nil, err
	} else {
		var ret, errResp = new(GetAudioPlayResponse), new(ErrorResponse)
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
