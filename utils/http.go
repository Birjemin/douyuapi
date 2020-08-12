package utils

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// HTTPClient ...
type HTTPClient struct {
	Client   *http.Client
	Response *http.Response
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// HTTPGet ...
func (c *HTTPClient) HTTPGet(url string, params map[string]string) (err error) {
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err == nil {
		req.URL.RawQuery = HTTPQueryBuild(params)
		c.Response, err = c.Client.Do(req)
	}
	return
}

// HTTPPost ...
func (c *HTTPClient) HTTPPost(url string, params map[string]string) error {
	var query = HTTPQueryBuild(params)
	return c.doPostRequest(url, query, "application/x-www-form-urlencoded;charset=UTF-8")
}

// HTTPPostJSON ...
func (c *HTTPClient) HTTPPostJSON(url, jsonStr string) error {
	return c.doPostRequest(url, jsonStr, "application/json;charset=UTF-8")
}

// doPostRequest ...
func (c *HTTPClient) doPostRequest(url, str, contentType string) (err error) {
	var req *http.Request
	if req, err = http.NewRequest("POST", url, strings.NewReader(str)); err != nil {
		return errors.New("sending http request error")
	}
	req.Header.Set("Content-Type", contentType)
	c.Response, err = c.Client.Do(req)
	return
}

// GetResponseJSON ...
func (c *HTTPClient) GetResponseJSON(response interface{}, errorResponse interface{}) error {
	if c.Response.Body == nil {
		return errors.New("http request response body is empty")
	}
	defer c.Response.Body.Close()

	data, err := c.GetResponseByte()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return json.Unmarshal(data, &errorResponse)
	}
	return nil
}

// GetResponseByte ...
func (c *HTTPClient) GetResponseByte() (body []byte, err error) {
	if c.Response.Body == nil {
		return []byte{}, errors.New("http request response body is empty")
	}
	defer c.Response.Body.Close()
	return ioutil.ReadAll(c.Response.Body)
}

// HTTPQueryBuild ...
func HTTPQueryBuild(params map[string]string) string {
	var query = make(url.Values)
	for k, v := range params {
		query.Add(k, v)
	}
	return query.Encode()
}
