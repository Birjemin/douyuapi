package douyuapi

import (
	"errors"
	"fmt"
	"github.com/birjemin/douyuapi/utils"
)

const DouYuDomain = "https://openapi.douyu.com"

// base client
type BaseClient struct {
	Client *utils.HttpClient
	Secret string
	AID    string
}

// Response
type ErrorResponse struct {
	Code int
	Msg  string
}

// GetErrResponse
func (e *ErrorResponse) GetErrResponse() error {
	if e.Code != 0 {
		return errors.New(fmt.Sprintf("%d|%s", e.Code, e.Msg))
	}
	return nil
}
