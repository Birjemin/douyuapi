package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const getRoomInfoURI = "/api/thirdPart/getRoomInfo"

// GetRoomInfo ...
type GetRoomInfo struct {
	BaseClient
	Token string
}

// GetRoomInfoResponse ...
type GetRoomInfoResponse struct {
	ErrorResponse
	Data struct {
		RID        int    `json:"rid"`
		RoomSrc    string `json:"room_src"`
		RoomSrcMax string `json:"room_src_max"`
		RoomName   string `json:"room_name"`
		Hn         int    `json:"hn"`
		Nickname   string `json:"nickname"`
		Avatar     string `json:"avatar"`
		Cid1       int    `json:"cid1"`
		Cname1     string `json:"cname1"`
		Cid2       int    `json:"cid2"`
		Cname2     string `json:"cname2"`
		Cid3       int    `json:"cid3"`
		Cname3     string `json:"cname3"`
		ShowStatus int    `json:"show_status"`
		ShowTime   int    `json:"show_time"`
		Unuid      string `json:"unuid"`
		RoomNotice string `json:"room_notice"`
		IsVertical int    `json:"is_vertical"`
		Fans       int    `json:"fans"`
	} `json:"data"`
}

// Handle ...
func (p *GetRoomInfo) Handle(postJSON, timestamp string) (*GetRoomInfoResponse, error) {
	return p.do(DouYuDomain+getRoomInfoURI, postJSON, timestamp)
}

// do
func (p *GetRoomInfo) do(url, postJSON, timestamp string) (*GetRoomInfoResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, getRoomInfoURI, params)

	url += "?" + utils.HTTPQueryBuild(params)

	if err := p.Client.HTTPPostJSON(url, postJSON); err != nil {
		return nil, err
	}
	var ret, errResp = new(GetRoomInfoResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJSON(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
