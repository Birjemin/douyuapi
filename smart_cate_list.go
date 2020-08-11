package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const smartCateListUri = "/api/thirdPart/smartCateList"

// SmartCateList
type SmartCateList struct {
	BaseClient
	Token string
}

// SmartCateListResponse
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
	} `json:"data;omitempty"`
}

// Handle
func (p *SmartCateList) Handle(postJson, timestamp string) (*SmartCateListResponse, error) {
	return p.do(DouYuDomain+smartCateListUri, postJson, timestamp)
}

// do
func (p *SmartCateList) do(url, postJson, timestamp string) (*SmartCateListResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, smartCateListUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJson); err != nil {
		return nil, err
	} else {
		var ret = new(SmartCateListResponse)
		if err := p.Client.GetResponseJson(ret); err != nil {
			return nil, err
		} else {
			return ret, nil
		}
	}
}
