package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// HttpClient
type HttpClient struct {
	Client      *http.Client
	Response    *http.Response
}

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
	if req, err = http.NewRequest("POST", url, strings.NewReader(str)); err == nil {
		req.Header.Set("Content-Type", contentType)
		c.Response, err = c.Client.Do(req)
	}
	return
}

// GetResponseJson
func (c *HttpClient) GetResponseJson(response interface{}) error {
	defer c.Response.Body.Close()
	return json.NewDecoder(c.Response.Body).Decode(response)
}

// GetResponseByte
func (c *HttpClient) GetResponseByte() []byte {
	var body []byte
	var err error

	defer c.Response.Body.Close()
	if body, err = ioutil.ReadAll(c.Response.Body); err != nil {
		log.Println("[http]method:GetResponseByte, body err: ", err)
	}
	return body
}

// HttpQueryBuild
func HttpQueryBuild(params map[string]string) string {
	var query = make(url.Values)
	for k, v := range params {
		query.Add(k, v)
	}
	return query.Encode()
}
