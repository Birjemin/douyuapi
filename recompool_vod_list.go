package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const recompoolVodListURI = "/api/thirdPart/video/recompoolVodList"

// RecompoolVodList ...
type RecompoolVodList struct {
	BaseClient
	Token string
}

// RecompoolVodListResponse ...
type RecompoolVodListResponse struct {
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
func (p *RecompoolVodList) Handle(postJSON, timestamp string) (*RecompoolVodListResponse, error) {
	return p.do(DouYuDomain+recompoolVodListURI, postJSON, timestamp)
}

// do
func (p *RecompoolVodList) do(url, postJSON, timestamp string) (*RecompoolVodListResponse, error) {
	var params = map[string]string{
		"aid":   p.AID,
		"time":  timestamp,
		"token": p.Token,
	}
	params["auth"] = GetSign(p.Secret, recompoolVodListURI, params)

	url += "?" + utils.HTTPQueryBuild(params)

	if err := p.Client.HTTPPostJSON(url, postJSON); err != nil {
		return nil, err
	}
	var ret, errResp = new(RecompoolVodListResponse), new(ErrorResponse)
	if err := p.Client.GetResponseJSON(ret, errResp); err != nil {
		return nil, err
	}
	if errResp.Code != 0 {
		ret.Code = errResp.Code
		ret.Msg = errResp.Msg
	}
	return ret, nil
}
