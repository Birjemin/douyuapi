package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
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

	c := &HttpClient{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	err := c.HttpGet(ts.URL, map[string]string{"param": "hello get", "param1": "yo~"})
	if err != nil {
		t.Fatal(err)
	}
	ast.Equal(c.GetResponseByte(), []byte("hello get"))
}

// TestGet
func TestGetJson(t *testing.T) {
	ast := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "%s", r.URL.Query()["param"][0])
	}))
	defer ts.Close()

	c := &HttpClient{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	err := c.HttpGet(ts.URL, map[string]string{"param": "{\"code\":1,\"msg\":\"ddd\"}", "param1": "yo~"})
	if err != nil {
		t.Fatal(err)
	}

	type JsonResponse struct {
		Code int
		Msg  string
	}
	var resp = new(JsonResponse)
	err = c.GetResponseJson(resp)
	if err != nil {
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

	c := &HttpClient{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	err := c.HttpPost(ts.URL, map[string]string{"param": "hello post", "param1": "yo~"})
	if err != nil {
		t.Fatal(err)
	}
	ast.Equal(c.GetResponseByte(), []byte("hello post"))
}

// TestPostJson
func TestPostJson(t *testing.T) {
	ast := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		_, _ = fmt.Fprintf(w, "%s", string(body))
	}))

	defer ts.Close()

	c := &HttpClient{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	if err := c.HttpPostJson(ts.URL, "{\"code\":1,\"msg\":\"ddd\"}"); err != nil {
		t.Fatal(err)
	} else {

		type JsonResponse struct {
			Code int
			Msg  string
		}
		var resp = new(JsonResponse)

		if err = c.GetResponseJson(resp); err != nil {
			t.Fatal(err)
		}
		ast.Equal(1, resp.Code)
		ast.Equal("ddd", resp.Msg)
	}
}

// TestHttpQueryBuild
func TestHttpQueryBuild(t *testing.T) {
	ast := assert.New(t)

	ret := HttpQueryBuild(map[string]string{"b": "11", "a": "22"})
	log.Println(ret)
	ast.Equal(ret, "a=22&b=11")
}
