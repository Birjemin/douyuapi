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
func (t *OuterChat) Handle(chat, timestamp string) (*OuterChatResponse, error) {
	return t.do(DouYuDomain+outerChatUri, chat, timestamp)
}

// do
func (t *OuterChat) do(url, chat, timestamp string) (*OuterChatResponse, error) {
	var params = map[string]string{
		"aid":   t.AID,
		"time":  timestamp,
		"token": t.Token,
	}
	params["auth"] = GetSign(t.Secret, outerChatUri, params)

	url += "?" + utils.HttpQueryBuild(params)
	if err := t.Client.HttpPostJson(url, chat); err != nil {
		return nil, err
	} else {
		var ret = new(OuterChatResponse)
		if err = t.Client.GetResponseJson(ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
}
