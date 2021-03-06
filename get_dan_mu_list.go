package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const getDanMuListURI = "/api/thirdPart/getDanMuList"

// GetDanMuList ...
type GetDanMuList struct {
	BaseClient
	Token string
}

// GetDanMuListResponse ...
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
	} `json:"data"`
}

// Handle ...
func (p *GetDanMuList) Handle(postJSON, timestamp string) (*GetDanMuListResponse, error) {
	return p.do(DouYuDomain+getDanMuListURI, postJSON, timestamp)
}

// do
func (p *GetDanMuList) do(url, postJSON, timestamp string) (*GetDanMuListResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, getDanMuListURI, params)

	url += "?" + utils.HTTPQueryBuild(params)
	if err := p.Client.HTTPPostJSON(url, postJSON); err != nil {
		return nil, err
	}
	var ret, errResp = new(GetDanMuListResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJSON(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
