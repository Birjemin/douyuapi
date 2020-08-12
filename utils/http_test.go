package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestGet
func TestGetString(t *testing.T) {
	ast := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "%s", r.URL.Query()["param"][0])
	}))
	defer ts.Close()

	c := &HTTPClient{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	if err := c.HTTPGet(ts.URL, map[string]string{"param": "hello get", "param1": "yo~"}); err != nil {
		t.Fatal(err)
	}
	if ret, err := c.GetResponseByte(); err != nil {
		t.Fatal(err)
	} else {
		ast.Equal(ret, []byte("hello get"))
	}
}

// TestGet
func TestGetJson(t *testing.T) {
	ast := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "%s", r.URL.Query()["param"][0])
	}))
	defer ts.Close()

	c := &HTTPClient{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	if err := c.HTTPGet(ts.URL, map[string]string{"param": "{\"code\":1,\"msg\":\"ddd\"}", "param1": "yo~"}); err != nil {
		t.Fatal(err)
	}

	type JSONResponse struct {
		Code int
		Msg  string
	}
	var resp = new(JSONResponse)
	if err := c.GetResponseJSON(resp, resp); err != nil {
		t.Fatal(err)
	}
	ast.Equal(1, resp.Code)
	ast.Equal("ddd", resp.Msg)
}

// TestPost
func TestPostString(t *testing.T) {
	ast := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "%s", r.FormValue("param"))
	}))
	defer ts.Close()

	c := &HTTPClient{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	if err := c.HTTPPost(ts.URL, map[string]string{"param": "hello post", "param1": "yo~"}); err != nil {
		t.Fatal(err)
	}

	if ret, err := c.GetResponseByte(); err != nil {
		t.Fatal(err)
	} else {
		ast.Equal(ret, []byte("hello post"))
	}
}

// TestPostJson
func TestPostJson(t *testing.T) {
	ast := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		_, _ = fmt.Fprintf(w, "%s", string(body))
	}))

	defer ts.Close()

	c := &HTTPClient{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	if err := c.HTTPPostJSON(ts.URL, "{\"code\":1,\"msg\":\"ddd\"}"); err != nil {
		t.Fatal(err)
	} else {

		type JSONResponse struct {
			Code int
			Msg  string
		}
		var resp = new(JSONResponse)

		if err = c.GetResponseJSON(resp, resp); err != nil {
			t.Fatal(err)
		}
		ast.Equal(1, resp.Code)
		ast.Equal("ddd", resp.Msg)
	}
}

// TestHTTPQueryBuild
func TestHTTPQueryBuild(t *testing.T) {
	ast := assert.New(t)

	ret := HTTPQueryBuild(map[string]string{"b": "11", "a": "22"})
	ast.Equal(ret, "a=22&b=11")
}
