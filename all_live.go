package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const allLiveURI = "/api/thirdPart/allLive"

// AllLive ...
type AllLive struct {
	BaseClient
	Token string
}

// AllLiveResponse ...
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
	} `json:"data"`
}

// Handle ...
func (p *AllLive) Handle(postJSON, timestamp string) (*AllLiveResponse, error) {
	return p.do(DouYuDomain+liveURI, postJSON, timestamp)
}

// do
func (p *AllLive) do(url, postJSON, timestamp string) (*AllLiveResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, liveURI, params)

	url += "?" + utils.HTTPQueryBuild(params)

	if err := p.Client.HTTPPostJSON(url, postJSON); err != nil {
		return nil, err
	}
	var ret, errResp = new(AllLiveResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJSON(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
