package douyuapi

import (
	"fmt"
	"github.com/birjemin/douyuapi/utils"
)

// DouYuDomain ...
const DouYuDomain = "https://openapi.douyu.com"

// BaseClient base client
type BaseClient struct {
	Client *utils.HttpClient
	Secret string
	AID    string
}

// ErrorResponse ...
type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// GetErrResponse ...
func (e *ErrorResponse) GetErrResponse() error {
	if e.Code != 0 {
		return fmt.Errorf("%d|%s", e.Code, e.Msg)
	}
	return nil
}
