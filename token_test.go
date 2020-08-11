package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
	"github.com/spf13/cast"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestToken
func TestToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != tokenUri {
			t.Fatalf("path is invalid: %s, %s'", tokenUri, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"aid", "auth", "time"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("param %v can not be empty", v)
			}
		}

		w.WriteHeader(http.StatusOK)

		var raw string
		if r.Form.Get("time") == "1200" {
			raw = `{"code": 1001, "msg":"", "data": ""}`
		} else {
			raw = `{"code": 0, "msg":"", "data": {"token":"ACCESS_TOKEN","expire":7200}}`
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

	token := &Token{
		BaseClient: BaseClient{
			Client: httpClient,
			Secret: "test-secret",
			AID:    "test-aid",
		},
	}

	timestamp := cast.ToString(utils.GetCurrTime())
	if ret, err := token.do(ts.URL+tokenUri, timestamp); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 0 {
			t.Fatal("msg: " + ret.Msg)
		}
		if ret.Data.Token != "ACCESS_TOKEN" {
			t.Fatal("get token failed")
		}
	}

	if ret, err := token.do(ts.URL+tokenUri, "1200"); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 1001 {
			t.Fatal("msg: " + ret.Msg)
		}
	}
}
