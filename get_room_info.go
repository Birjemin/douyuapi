package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const roomInfoUri = "/api/thirdPart/getRoomInfo"

// RoomInfo
type RoomInfo struct {
	BaseClient
	Token string
}

// RoomInfoResponse
type RoomInfoResponse struct {
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

// Handle
func (r *RoomInfo) Handle(postJson, timestamp string) (*RoomInfoResponse, error) {
	return r.do(DouYuDomain+roomInfoUri, postJson, timestamp)
}

// do
func (r *RoomInfo) do(url, postJson, timestamp string) (*RoomInfoResponse, error) {
	var params = map[string]string{
		"aid":   r.AID,
		"time":  timestamp,
		"token": r.Token,
	}
	params["auth"] = GetSign(r.Secret, roomInfoUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := r.Client.HttpPostJson(url, postJson); err != nil {
		return nil, err
	} else {
		var ret = new(RoomInfoResponse)
		if err = r.Client.GetResponseJson(ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
}
