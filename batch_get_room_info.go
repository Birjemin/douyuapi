package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const batchRoomInfoURI = "/api/thirdPart/batchGetRoomInfo"

// BatchRoomInfo ...
type BatchRoomInfo struct {
	BaseClient
	Token string
}

// BatchRoomInfoResponse ...
type BatchRoomInfoResponse struct {
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
func (p *BatchRoomInfo) Handle(postJSON, timestamp string) (*BatchRoomInfoResponse, error) {
	return p.do(DouYuDomain+batchRoomInfoURI, postJSON, timestamp)
}

// do
func (p *BatchRoomInfo) do(url, postJSON, timestamp string) (*BatchRoomInfoResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, batchRoomInfoURI, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJSON); err != nil {
		return nil, err
	}
	var ret, errResp = new(BatchRoomInfoResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJson(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
