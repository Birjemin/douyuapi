package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const cid2InfoUri = "/api/thirdPart/getCid2Info"

// Cid2Info
type Cid2Info struct {
	BaseClient
	Token string
}

// Cid2InfoResponse
type Cid2InfoResponse struct {
	ErrorResponse
	Data []struct {
		CID2   int    `json:"cid2"`
		CName2 string `json:"cname2"`
		PicUrl string `json:"pic_url"`
	} `json:"data;omitempty"`
}

// Handle
func (r *Cid2Info) Handle(timestamp string) (*Cid2InfoResponse, error) {
	return r.do(DouYuDomain+cid2InfoUri, timestamp)
}

// do
func (r *Cid2Info) do(url, timestamp string) (*Cid2InfoResponse, error) {
	var params = map[string]string{
		"aid":   r.AID,
		"time":  timestamp,
		"token": r.Token,
	}
	params["auth"] = GetSign(r.Secret, cid2InfoUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := r.Client.HttpPostJson(url, ""); err != nil {
		return nil, err
	} else {
		var ret = new(Cid2InfoResponse)
		if err = r.Client.GetResponseJson(ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
}
