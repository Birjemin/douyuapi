package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const allLiveUri = "/api/thirdPart/allLive"

// AllLive
type AllLive struct {
	BaseClient
	Token string
}

// AllLiveResponse
type AllLiveResponse struct {
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
	} `json:"data;omitempty"`
}

// Handle
func (p *AllLive) Handle(postJson, timestamp string) (*AllLiveResponse, error) {
	return p.do(DouYuDomain+liveUri, postJson, timestamp)
}

// do
func (p *AllLive) do(url, postJson, timestamp string) (*AllLiveResponse, error) {
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
		var ret = new(AllLiveResponse)
		if err := p.Client.GetResponseJson(ret); err != nil {
			return nil, err
		} else {
			return ret, nil
		}

	}
}
