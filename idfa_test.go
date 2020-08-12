package douyuapi

import (
	"errors"
	"fmt"
	"github.com/birjemin/douyuapi/utils"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestIDFA
func TestIDFA(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != idfaURI {
			t.Fatalf("path is invalid: %s, %s'", idfaURI, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"aid", "auth", "time", "token"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("param %v can not be empty", v)
			}
		}

		body, _ := ioutil.ReadAll(r.Body)
		if string(body) == "" {
			t.Fatal("body is empty")
		}

		w.WriteHeader(http.StatusOK)

		var raw string
		if r.Form.Get("time") == "100" {
			raw = `{"code": 100, "msg":"", "data": ""}`
		} else {
			raw = `{"code": 0, "msg":"ok", "data": {"080006E2-5666-49C1-8786-3FD9FC77DC0A":1}}`
		}

		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))

	defer ts.Close()

	httpClient := &utils.HttpClient{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	idfaInfo := &IDFAInfo{
		BaseClient: BaseClient{
			Client: httpClient,
			Secret: "test-secret",
			AID:    "test-aid",
		},
		Token: "test-token",
	}

	timestamp := utils.GetCurrTime()

	chat := fmt.Sprintf(
		`{"idfa":"%s,%s"}`,
		"080006E2-5666-49C1-8786-3FD9FC77DC0A",
		"080006E2-5666-49C1-8786-3FD9FC77DC0B",
	)

	if ret, err := idfaInfo.do(ts.URL+idfaURI, chat, cast.ToString(timestamp)); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 0 {
			t.Error(errors.New("msg: " + ret.Msg))
		}
	}

	if ret, err := idfaInfo.do(ts.URL+idfaURI, chat, "100"); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 100 {
			t.Error(errors.New("msg: " + ret.Msg))
		}
	}
}
