package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const videoCateVodListUri = "/api/thirdPart/video/video_cate_list.go"

// VideoCateVodList
type VideoCateVodList struct {
	BaseClient
	Token string
}

// VideoCateVodListResponse
type VideoCateVodListResponse struct {
	ErrorResponse
	Data []struct {
		Cid2   int    `json:"cid2"`
		Cname2 string `json:"cname2"`
	} `json:"data"`
}

// Handle
func (p *VideoCateVodList) Handle(postJson, timestamp string) (*VideoCateListResponse, error) {
	return p.do(DouYuDomain+videoCateVodListUri, postJson, timestamp)
}

// do
func (p *VideoCateVodList) do(url, postJson, timestamp string) (*VideoCateListResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, videoCateVodListUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJson); err != nil {
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
