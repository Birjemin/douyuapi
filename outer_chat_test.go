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

// TestOuterChat
func TestOuterChat(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != outerChatURI {
			t.Fatalf("path is invalid: %s, %s'", outerChatURI, path)
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
			raw = `{"code": 0, "msg":"ok", "data": ""}`
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

	outChat := &OuterChat{
		BaseClient: BaseClient{
			Client: httpClient,
			Secret: "test-secret",
			AID:    "test-aid",
		},
		Token: "test-token",
	}

	timestamp := utils.GetCurrTime()

	chat := fmt.Sprintf(`{"chat":[{"room_id":%d,"uid":%d,"nick_name":"%s","content":"%s","timestamp":%d,"ip":"%s"}]}`, 1000, 1000, "test", "hello", timestamp, "10.0.0.0")

	if ret, err := outChat.do(ts.URL+outerChatURI, chat, cast.ToString(timestamp)); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 0 {
			t.Error(errors.New("msg: " + ret.Msg))
		}
	}

	if ret, err := outChat.do(ts.URL+outerChatURI, chat, "100"); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 100 {
			t.Error(errors.New("msg: " + ret.Msg))
		}
	}
}
