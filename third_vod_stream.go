package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const thirdVodStreamUri = "/api/thirdPart/video/thirdVodStream"

// ThirdVodStreamList
type ThirdVodStreamList struct {
	BaseClient
	Token string
}

// ThirdVodStreamListResponse
type ThirdVodStreamListResponse struct {
	ErrorResponse
	Data struct {
		Timestamp int    `json:"timestamp"`
		Normal    string `json:"normal"`
		High      string `json:"high"`
	} `json:"data"`
}

// Handle
func (p *ThirdVodStreamList) Handle(postJson, timestamp string) (*ThirdVodStreamListResponse, error) {
	return p.do(DouYuDomain+thirdVodStreamUri, postJson, timestamp)
}

// do
func (p *ThirdVodStreamList) do(url, postJson, timestamp string) (*ThirdVodStreamListResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, thirdVodStreamUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJson); err != nil {
		return nil, err
	} else {
		var ret, errResp = new(ThirdVodStreamListResponse), new(ErrorResponse)
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
