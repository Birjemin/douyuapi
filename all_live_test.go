package douyuapi

import (
	"errors"
	"github.com/birjemin/douyuapi/utils"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestAllLive ...
func TestAllLive(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != allLiveURI {
			t.Fatalf("path is invalid: %s, %s'", allLiveURI, path)
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
			raw = `{"code":0,"msg":"","data":[{"rid":1,"room_src":"","room_src_max":"","room_name":"","hn":1,"nickname":"","avatar":"","cid1":1,"cname1":"","cid2":1,"cname2":"","cid3":1,"cname3":"","show_status":1,"show_time":0,"unuid":""}]}`
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

	allLive := &AllLive{
		BaseClient: BaseClient{
			Client: httpClient,
			Secret: "test-secret",
			AID:    "test-aid",
		},
		Token: "test-token",
	}

	timestamp := utils.GetCurrTime()

	msg := `{"cid_type":1,"cid":1}`

	if ret, err := allLive.do(ts.URL+allLiveURI, msg, cast.ToString(timestamp)); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 0 || len(ret.Data) != 1 {
			t.Error(errors.New("msg: " + ret.Msg))
		}
	}

	if ret, err := allLive.do(ts.URL+allLiveURI, msg, "100"); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 100 {
			t.Error(errors.New("msg: " + ret.Msg))
		}
	}
}
