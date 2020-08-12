package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const thirdVodStreamURI = "/api/thirdPart/video/thirdVodStream"

// ThirdVodStreamList ...
type ThirdVodStreamList struct {
	BaseClient
	Token string
}

// ThirdVodStreamListResponse ...
type ThirdVodStreamListResponse struct {
	ErrorResponse
	Data struct {
		Timestamp int    `json:"timestamp"`
		Normal    string `json:"normal"`
		High      string `json:"high"`
	} `json:"data"`
}

// Handle ...
func (p *ThirdVodStreamList) Handle(postJSON, timestamp string) (*ThirdVodStreamListResponse, error) {
	return p.do(DouYuDomain+thirdVodStreamURI, postJSON, timestamp)
}

// do
func (p *ThirdVodStreamList) do(url, postJSON, timestamp string) (*ThirdVodStreamListResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, thirdVodStreamURI, params)

	url += "?" + utils.HTTPQueryBuild(params)

	if err := p.Client.HTTPPostJSON(url, postJSON); err != nil {
		return nil, err
	}
	var ret, errResp = new(ThirdVodStreamListResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJSON(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
