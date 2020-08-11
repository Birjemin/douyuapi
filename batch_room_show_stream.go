package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const batchRoomShowStreamUri = "/api/thirdPart/batchRoomShowStream"

// BatchRoomShowStream
type BatchRoomShowStream struct {
	BaseClient
	Token string
}

// BatchRoomShowStreamResponse
type BatchRoomShowStreamResponse struct {
	ErrorResponse
	Data struct {
		RtmpID int    `json:"rtmp_id"`
		RID    int    `json:"rid"`
		Nfv    string `json:"nfv"`
	} `json:"data;omitempty"`
}

// Handle
func (p *BatchRoomShowStream) Handle(postJson, timestamp string) (*BatchRoomShowStreamResponse, error) {
	return p.do(DouYuDomain+batchRoomShowStreamUri, postJson, timestamp)
}

// do
func (p *BatchRoomShowStream) do(url, postJson, timestamp string) (*BatchRoomShowStreamResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, batchRoomShowStreamUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJson); err != nil {
		return nil, err
	} else {
		var ret, errResp = new(BatchRoomShowStreamResponse), new(ErrorResponse)
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
