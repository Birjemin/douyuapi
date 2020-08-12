package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const outerChatURI = "/api/thirdPart/outerChat"

// OuterChat ...
type OuterChat struct {
	BaseClient
	Token string
}

// OuterChatResponse ...
type OuterChatResponse struct {
	ErrorResponse
	Data string `json:"data"`
}

// Handle ...
func (p *OuterChat) Handle(chat, timestamp string) (*OuterChatResponse, error) {
	return p.do(DouYuDomain+outerChatURI, chat, timestamp)
}

// do
func (p *OuterChat) do(url, chat, timestamp string) (*OuterChatResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, outerChatURI, params)

	url += "?" + utils.HTTPQueryBuild(params)
	if err := p.Client.HTTPPostJSON(url, chat); err != nil {
		return nil, err
	}
	var ret, errResp = new(OuterChatResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJSON(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
