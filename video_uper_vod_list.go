package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const videoUperVodListUri = "/api/thirdPart/video/uperVodList"

// VideoUperVodList
type VideoUperVodList struct {
	BaseClient
	Token string
}

// VideoUperVodListResponse
type VideoUperVodListResponse struct {
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
func (p *VideoUperVodList) Handle(postJson, timestamp string) (*VideoUperVodListResponse, error) {
	return p.do(DouYuDomain+videoUperVodListUri, postJson, timestamp)
}

// do
func (p *VideoUperVodList) do(url, postJson, timestamp string) (*VideoUperVodListResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, videoUperVodListUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJson); err != nil {
		return nil, err
	} else {
		var ret, errResp = new(VideoUperVodListResponse), new(ErrorResponse)
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
