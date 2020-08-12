package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const getAudioPlayURI = "/api/thirdPart/getAudioPlay"

// GetAudioPlay ...
type GetAudioPlay struct {
	BaseClient
	Token string
}

// GetAudioPlayResponse ...
type GetAudioPlayResponse struct {
	ErrorResponse
	Data []struct {
		RID      int    `json:"rid"`
		UID      int    `json:"uid"`
		RtmpCdn  string `json:"rtmpCdn"`
		RtmpLive string `json:"rtmpLive"`
		RtmpURL  string `json:"rtmpURL"`
	} `json:"data"`
}

// Handle ...
func (p *GetAudioPlay) Handle(postJSON, timestamp string) (*GetAudioPlayResponse, error) {
	return p.do(DouYuDomain+getAudioPlayURI, postJSON, timestamp)
}

// do
func (p *GetAudioPlay) do(url, postJSON, timestamp string) (*GetAudioPlayResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, getAudioPlayURI, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJSON); err != nil {
		return nil, err
	}
	var ret, errResp = new(GetAudioPlayResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJson(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
