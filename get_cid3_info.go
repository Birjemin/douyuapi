package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const Cid3InfoUri = "/api/thirdPart/getCid3Info"

// Cid3Info
type Cid3Info struct {
	BaseClient
	Token string
}

// Cid3InfoResponse
type Cid3InfoResponse struct {
	ErrorResponse
	Data []struct {
		CID3   int    `json:"cid3"`
		CName3 string `json:"cname3"`
	} `json:"data"`
}

// Handle
func (r *Cid3Info) Handle(timestamp string) (*Cid3InfoResponse, error) {
	return r.do(DouYuDomain+Cid3InfoUri, timestamp)
}

// do
func (r *Cid3Info) do(url string, timestamp string) (*Cid3InfoResponse, error) {
	var params = map[string]string{
		"aid":   r.AID,
		"time":  timestamp,
		"token": r.Token,
	}
	params["auth"] = GetSign(r.Secret, Cid3InfoUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := r.Client.HttpPostJson(url, ""); err != nil {
		return nil, err
	} else {
		var ret = new(Cid3InfoResponse)
		if err = r.Client.GetResponseJson(ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
}
