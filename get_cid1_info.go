package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
)

const cid1InfoUri = "/api/thirdPart/getCid1Info"

// Cid1Info
type Cid1Info struct {
	BaseClient
	Token string
}

// Cid1InfoResponse
type Cid1InfoResponse struct {
	ErrorResponse
	Data []struct {
		CID1   int    `json:"cid1"`
		CName1 string `json:"cname1"`
	} `json:"data;omitempty"`
}

// Handle
func (r *Cid1Info) Handle(timestamp string) (*Cid1InfoResponse, error) {
	return r.do(DouYuDomain+cid1InfoUri, timestamp)
}

// do
func (r *Cid1Info) do(url, timestamp string) (*Cid1InfoResponse, error) {
	var params = map[string]string{
		"aid":   r.AID,
		"time":  timestamp,
		"token": r.Token,
	}
	params["auth"] = GetSign(r.Secret, cid1InfoUri, params)

	url += "?" + utils.HttpQueryBuild(params)

	if err := r.Client.HttpPostJson(url, ""); err != nil {
		return nil, err
	} else {
		var ret = new(Cid1InfoResponse)
		if err = r.Client.GetResponseJson(ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
}
