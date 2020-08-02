package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const BatchRoomInfoUri = "/api/thirdPart/batchGetRoomInfo"

// BatchRoomInfo
type BatchRoomInfo struct {
	BaseClient
	Token string
}

// BatchRoomInfoResponse
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

// Handle
func (r *BatchRoomInfo) Handle(postJson, timestamp string) (*BatchRoomInfoResponse, error) {
	return r.do(DouYuDomain+BatchRoomInfoUri, postJson, timestamp)
}

// do
func (r *BatchRoomInfo) do(url string, postJson string, timestamp string) (*BatchRoomInfoResponse, error) {
	var params = map[string]string{
		"aid":   r.AID,
		"time":  timestamp,
		"token": r.Token,
	}
	params["auth"] = GetSign(r.Secret, BatchRoomInfoUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := r.Client.HttpPostJson(url, postJson); err != nil {
		return nil, err
	} else {
		var ret = new(BatchRoomInfoResponse)
		if err = r.Client.GetResponseJson(ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
}
