package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const batchGetRoomInfoURI = "/api/thirdPart/batchGetRoomInfo"

// BatchGetRoomInfo ...
type BatchGetRoomInfo struct {
	BaseClient
	Token string
}

// BatchGetRoomInfoResponse ...
type BatchGetRoomInfoResponse struct {
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
		ShowStatus int    `json:"show_status"`
		RoomNotice string `json:"room_notice"`
		IsVertical int    `json:"is_vertical"`
	} `json:"data"`
}

// Handle ...
func (p *BatchGetRoomInfo) Handle(postJSON, timestamp string) (*BatchGetRoomInfoResponse, error) {
	return p.do(DouYuDomain+batchGetRoomInfoURI, postJSON, timestamp)
}

// do
func (p *BatchGetRoomInfo) do(url, postJSON, timestamp string) (*BatchGetRoomInfoResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, batchGetRoomInfoURI, params)

	url += "?" + utils.HTTPQueryBuild(params)

	if err := p.Client.HTTPPostJSON(url, postJSON); err != nil {
		return nil, err
	}
	var ret, errResp = new(BatchGetRoomInfoResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJSON(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
