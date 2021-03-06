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
		if path != tokenURI {
			t.Fatalf("path is invalid: %s, %s'", tokenURI, path)
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
		if r.Form.Get("time") == "100" {
			raw = `{"code": 100, "msg":"", "data": ""}`
		} else {
			raw = `{"code": 0, "msg":"", "data": {"token":"ACCESS_TOKEN","expire":7200}}`
		}

		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))

	defer ts.Close()

	httpClient := &utils.HTTPClient{
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
	if ret, err := token.do(ts.URL+tokenURI, timestamp); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 0 {
			t.Fatal("msg: " + ret.Msg)
		}
		if ret.Data.Token != "ACCESS_TOKEN" {
			t.Fatal("get token failed")
		}
	}

	if ret, err := token.do(ts.URL+tokenURI, "100"); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 100 {
			t.Fatal("msg: " + ret.Msg)
		}
	}
}
