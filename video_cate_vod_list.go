package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const videoCateVodListURI = "/api/thirdPart/video/cateVodList"

// VideoCateVodList ...
type VideoCateVodList struct {
	BaseClient
	Token string
}

// VideoCateVodListResponse ...
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
		VideoURL           string `json:"video_url"`
		H5VideoURL         string `json:"h5_video_url"`
		ShareURL           string `json:"share_url"`
	} `json:"data"`
}

// Handle ...
func (p *VideoCateVodList) Handle(postJSON, timestamp string) (*VideoCateVodListResponse, error) {
	return p.do(DouYuDomain+videoCateVodListURI, postJSON, timestamp)
}

// do
func (p *VideoCateVodList) do(url, postJSON, timestamp string) (*VideoCateVodListResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, videoCateVodListURI, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := p.Client.HttpPostJson(url, postJSON); err != nil {
		return nil, err
	}
	var ret, errResp = new(VideoCateVodListResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJson(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
