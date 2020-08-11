package utils

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// HttpClient
type HttpClient struct {
	Client   *http.Client
	Response *http.Response
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// HttpGet
func (c *HttpClient) HttpGet(url string, params map[string]string) (err error) {
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err == nil {
		req.URL.RawQuery = HttpQueryBuild(params)
		c.Response, err = c.Client.Do(req)
	}
	return
}

// HttpPost
func (c *HttpClient) HttpPost(url string, params map[string]string) error {
	var query = HttpQueryBuild(params)
	return c.doPostRequest(url, query, "application/x-www-form-urlencoded;charset=UTF-8")
}

// HttpPost
func (c *HttpClient) HttpPostJson(url, jsonStr string) error {
	return c.doPostRequest(url, jsonStr, "application/json;charset=UTF-8")
}

// doPostRequest
func (c *HttpClient) doPostRequest(url, str, contentType string) (err error) {
	var req *http.Request
	if req, err = http.NewRequest("POST", url, strings.NewReader(str)); err != nil {
		return errors.New("sending http request error")
	} else {
		req.Header.Set("Content-Type", contentType)
		c.Response, err = c.Client.Do(req)
	}
	return
}

// GetResponseJson
func (c *HttpClient) GetResponseJson(response interface{}, errorResponse interface{}) error {
	if c.Response.Body == nil {
		return errors.New("http request response body is empty")
	}
	defer c.Response.Body.Close()

	if data, err := c.GetResponseByte(); err != nil {
		return err
	} else {
		if err := json.Unmarshal(data, &response); err != nil {
			return json.Unmarshal(data, &errorResponse)
		} else {
			return nil
		}
	}
}

// GetResponseByte
func (c *HttpClient) GetResponseByte() (body []byte, err error) {
	if c.Response.Body == nil {
		return []byte{}, errors.New("http request response body is empty")
	}
	defer c.Response.Body.Close()
	return ioutil.ReadAll(c.Response.Body)
}

// HttpQueryBuild
func HttpQueryBuild(params map[string]string) string {
	var query = make(url.Values)
	for k, v := range params {
		query.Add(k, v)
	}
	return query.Encode()
}
