package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const smartCateListURI = "/api/thirdPart/smartCateList"

// SmartCateList ...
type SmartCateList struct {
	BaseClient
	Token string
}

// SmartCateListResponse ...
type SmartCateListResponse struct {
	ErrorResponse
	Data []struct {
		Tag1ID    int    `json:"tag1_id"`
		Tag1Name  string `json:"tag1_name"`
		Tag1Icon  string `json:"tag1_icon"`
		Tag2ID    int    `json:"tag2_id"`
		Tag2Name  string `json:"tag2_name"`
		Tag2Icon  string `json:"tag2_icon"`
		RoomCount int    `json:"room_count"`
	} `json:"data"`
}

// Handle ...
func (p *SmartCateList) Handle(postJSON, timestamp string) (*SmartCateListResponse, error) {
	return p.do(DouYuDomain+smartCateListURI, postJSON, timestamp)
}

// do
func (p *SmartCateList) do(url, postJSON, timestamp string) (*SmartCateListResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, smartCateListURI, params)

	url += "?" + utils.HTTPQueryBuild(params)

	if err := p.Client.HTTPPostJSON(url, postJSON); err != nil {
		return nil, err
	}
	var ret, errResp = new(SmartCateListResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJSON(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
