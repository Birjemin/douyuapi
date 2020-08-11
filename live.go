package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const liveUri = "/api/thirdPart/live"

// Live
type Live struct {
	BaseClient
	Token string
}

// LiveResponse
type LiveResponse struct {
	ErrorResponse
	Data []struct {
		RID      int    `json:"rid"`
		RoomName string `json:"room_name"`
		RoomSrc  string `json:"room_src"`
		Hn       int    `json:"hn"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
		Cid1     int    `json:"cid1"`
		Cname1   string `json:"cname1"`
		Cid2     int    `json:"cid2"`
		Cname2   string `json:"cname2"`
		Cid3     int    `json:"cid3"`
		Cname3   string `json:"cname3"`
	} `json:"data"`
}

// Handle
func (p *Live) Handle(postJson, timestamp string) (*LiveResponse, error) {
	return p.do(DouYuDomain+liveUri, postJson, timestamp)
}

// do
func (p *Live) do(url, postJson, timestamp string) (*LiveResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, liveUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJson); err != nil {
		return nil, err
	} else {
		var ret, errResp = new(LiveResponse), new(ErrorResponse)
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
