package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const videoCateListUri = "/api/thirdPart/video/cateList"

// VideoCateList
type VideoCateList struct {
	BaseClient
	Token string
}

// VideoCateListResponse
type VideoCateListResponse struct {
	ErrorResponse
	Data []struct {
		Cid2     int    `json:"cid2"`
		Cname2   string `json:"cname2"`
	} `json:"data"`
}

// Handle
func (p *VideoCateList) Handle(timestamp string) (*VideoCateListResponse, error) {
	return p.do(DouYuDomain+videoCateListUri, timestamp)
}

// do
func (p *VideoCateList) do(url, timestamp string) (*VideoCateListResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, videoCateListUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, "{}"); err != nil {
		return nil, err
	} else {
		var ret = new(VideoCateListResponse)
		if err := p.Client.GetResponseJson(ret); err != nil {
			return nil, err
		} else {
			return ret, nil
		}

	}
}

