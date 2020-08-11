package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const idfaUri = "/api/thirdPart/idfa"

// IDFAInfo
type IDFAInfo struct {
	BaseClient
	Token string
}

// IDFAResponse
type IDFAResponse struct {
	ErrorResponse
	Data map[string]int `json:"data;omitempty"`
}

// Handle
func (r *IDFAInfo) Handle(idfa, timestamp string) (*IDFAResponse, error) {
	return r.do(DouYuDomain+idfaUri, idfa, timestamp)
}

// do
func (r *IDFAInfo) do(url, idfa, timestamp string) (*IDFAResponse, error) {
	var params = map[string]string{
		"aid":   r.AID,
		"time":  timestamp,
		"token": r.Token,
	}
	params["auth"] = GetSign(r.Secret, idfaUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := r.Client.HttpPostJson(url, idfa); err != nil {
		return nil, err
	} else {
		var ret = new(IDFAResponse)
		if err = r.Client.GetResponseJson(ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
}
