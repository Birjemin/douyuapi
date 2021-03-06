package douyuapi

import (
	"github.com/birjemin/douyuapi/utils"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestGetPlay
func TestGetPlay(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != getPlayURI {
			t.Fatalf("path is invalid: %s, %s'", getPlayURI, path)
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
			raw = `{"code":0,"msg":"","data":{"rid":288016,"room_name":"赛事直播间","live_url":"http+flv流地址","hls_url":"m3u8流地址","mix_url":"","rate_switch":0,"show_status":1,"hls_mul":{},"flv_mul":{}}}`
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

	play := &GetPlay{
		BaseClient: BaseClient{
			Client: httpClient,
			Secret: "test-secret",
			AID:    "test-aid",
		},
		Token: "test-token",
	}

	timestamp := utils.GetCurrTime()

	if ret, err := play.do(ts.URL+getPlayURI, `{"rid": 288016}`, cast.ToString(timestamp)); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 0 {
			t.Fatal("msg: " + ret.Msg)
		}
		if ret.Data.RID != 288016 {
			t.Fatal("get play failed")
		}
	}

	if ret, err := play.do(ts.URL+getPlayURI, `{"rid": 288016}`, "100"); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 100 {
			t.Fatal("msg: " + ret.Msg)
		}
	}
}
