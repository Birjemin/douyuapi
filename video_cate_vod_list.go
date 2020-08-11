package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const videoCateVodListUri = "/api/thirdPart/video/cateVodList"

// VideoCateVodList
type VideoCateVodList struct {
	BaseClient
	Token string
}

// VideoCateVodListResponse
type VideoCateVodListResponse struct {
	ErrorResponse
	Data []struct {
		HashID             string `json:"hash_id"`
		Cid1               int    `json:"cid1"`
		Cid2               int    `json:"cid2"`
		UpID               string `json:"up_id"`
		Nickname           string `json:"nickname"`
		IsVertical         int    `json:"is_vertical"`
		VideoTitle         string `json:"video_title"`
		VideoCover         string `json:"video_cover"`
		VideoVerticalCover string `json:"video_vertical_cover"`
		VideoDuration      int    `json:"video_duration"`
		ViewNum            int    `json:"view_num"`
		UTime              int    `json:"utime"`
		OwnerAvatar        string `json:"owner_avatar"`
		VideoCollectNum    int    `json:"video_collect_num"`
		VideoUpNum         int    `json:"video_up_num"`
		BarrageNum         int    `json:"barrage_num"`
		VideoUrl           string `json:"video_url"`
		H5VideoUrl         string `json:"h5_video_url"`
		ShareUrl           string `json:"share_url"`
	} `json:"data;omitempty"`
}

// Handle
func (p *VideoCateVodList) Handle(postJson, timestamp string) (*VideoCateVodListResponse, error) {
	return p.do(DouYuDomain+videoCateVodListUri, postJson, timestamp)
}

// do
func (p *VideoCateVodList) do(url, postJson, timestamp string) (*VideoCateVodListResponse, error) {
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
		var ret = new(VideoCateVodListResponse)
		if err := p.Client.GetResponseJson(ret); err != nil {
			return nil, err
		} else {
			return ret, nil
		}

	}
}
