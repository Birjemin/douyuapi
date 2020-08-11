package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const smartCateRoomListUri = "/api/thirdPart/smartCateRoomList"

// SmartCateRoomList
type SmartCateRoomList struct {
	BaseClient
	Token string
}

// SmartCateRoomListResponse
type SmartCateRoomListResponse struct {
	ErrorResponse
	Data []struct {
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
		RoomNotice string `json:"room_notice"`
		IsVertical int    `json:"is_vertical"`
		ShowStatus int    `json:"show_status"`
		Tag2ID     int    `json:"tag2_id"`
		Tag2Name   string `json:"tag2_name"`
		Tag2Icon   string `json:"tag2_icon"`
		Count      int    `json:"count"`
	} `json:"data;omitempty"`
}

// Handle
func (p *SmartCateRoomList) Handle(postJson, timestamp string) (*SmartCateRoomListResponse, error) {
	return p.do(DouYuDomain+smartCateRoomListUri, postJson, timestamp)
}

// SmartCateRoomList
func (p *SmartCateRoomList) do(url, postJson, timestamp string) (*SmartCateRoomListResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, smartCateRoomListUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJson); err != nil {
		return nil, err
	} else {
		var ret = new(SmartCateRoomListResponse)
		if err := p.Client.GetResponseJson(ret); err != nil {
			return nil, err
		} else {
			return ret, nil
		}
	}
}
