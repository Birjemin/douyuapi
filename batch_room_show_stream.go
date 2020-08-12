package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const batchRoomShowStreamURI = "/api/thirdPart/batchRoomShowStream"

// BatchRoomShowStream ...
type BatchRoomShowStream struct {
	BaseClient
	Token string
}

// BatchRoomShowStreamResponse ...
type BatchRoomShowStreamResponse struct {
	ErrorResponse
	Data struct {
		RtmpID int    `json:"rtmp_id"`
		RID    int    `json:"rid"`
		Nfv    string `json:"nfv"`
	} `json:"data"`
}

// Handle ...
func (p *BatchRoomShowStream) Handle(postJSON, timestamp string) (*BatchRoomShowStreamResponse, error) {
	return p.do(DouYuDomain+batchRoomShowStreamURI, postJSON, timestamp)
}

// do
func (p *BatchRoomShowStream) do(url, postJSON, timestamp string) (*BatchRoomShowStreamResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, batchRoomShowStreamURI, params)

	url += "?" + utils.HTTPQueryBuild(params)

	if err := p.Client.HTTPPostJSON(url, postJSON); err != nil {
		return nil, err
	}
	var ret, errResp = new(BatchRoomShowStreamResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJSON(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
