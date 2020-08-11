package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const outerChatUri = "/api/thirdPart/outerChat"

// Token
type OuterChat struct {
	BaseClient
	Token string
}

// OuterChatResponse
type OuterChatResponse struct {
	ErrorResponse
	Data string `json:"data"`
}

// Handle
func (p *OuterChat) Handle(chat, timestamp string) (*OuterChatResponse, error) {
	return p.do(DouYuDomain+outerChatUri, chat, timestamp)
}

// do
func (p *OuterChat) do(url, chat, timestamp string) (*OuterChatResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, outerChatUri, params)

	url += "?" + utils.HttpQueryBuild(params)
	if err := p.Client.HttpPostJson(url, chat); err != nil {
		return nil, err
	} else {
		var ret, errResp = new(OuterChatResponse), new(ErrorResponse)
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
