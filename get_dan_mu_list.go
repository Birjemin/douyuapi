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
func (t *GetDanMuList) Handle(postJson, timestamp string) (*GetDanMuListResponse, error) {
	return t.do(DouYuDomain+getDanMuListUri, postJson, timestamp)
}

// do
func (t *GetDanMuList) do(url, postJson, timestamp string) (*GetDanMuListResponse, error) {
	var params = map[string]string{
		"aid":   t.AID,
		"time":  timestamp,
		"token": t.Token,
	}
	params["auth"] = GetSign(t.Secret, getDanMuListUri, params)

	url += "?" + utils.HttpQueryBuild(params)
	if err := t.Client.HttpPostJson(url, postJson); err != nil {
		return nil, err
	} else {
		var ret = new(GetDanMuListResponse)
		if err = t.Client.GetResponseJson(ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
}
