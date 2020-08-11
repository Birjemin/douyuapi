package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const getDanMuListUri = "/api/thirdPart/getDanMuList"

// GetDanMuList
type GetDanMuList struct {
	BaseClient
	Token string
}

// GetDanMuListResponse
type GetDanMuListResponse struct {
	ErrorResponse
	Data struct {
		List []struct {
			RoomID    int    `json:"room_id"`
			UID       int    `json:"uid"`
			Nickname  string `json:"nickname"`
			Content   string `json:"content"`
			Timestamp int    `json:"timestamp"`
			IP        string `json:"ip"`
		} `json:"list"`
		Cnt         int `json:"cnt"`
		PageContext int `json:"page_context"`
	} `json:"data;omitempty"`
}

// Handle
func (p *GetDanMuList) Handle(postJson, timestamp string) (*GetDanMuListResponse, error) {
	return p.do(DouYuDomain+getDanMuListUri, postJson, timestamp)
}

// do
func (p *GetDanMuList) do(url, postJson, timestamp string) (*GetDanMuListResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, getDanMuListUri, params)

	url += "?" + utils.HttpQueryBuild(params)
	if err := p.Client.HttpPostJson(url, postJson); err != nil {
		return nil, err
	} else {
		var ret, errResp = new(GetDanMuListResponse), new(ErrorResponse)
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
