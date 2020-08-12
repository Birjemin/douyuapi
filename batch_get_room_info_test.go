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

// TestBatchGetRoomInfo
func TestBatchGetRoomInfo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != batchRoomInfoURI {
			t.Fatalf("path is invalid: %s, %s'", batchRoomInfoURI, path)
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
			raw = `{"code":0,"msg":"","data":[{"rid":688,"room_src":"room_src","room_src_max":"room_src_max","room_name":"房间","hn":6030,"nickname":"test","avatar":"a7e997_big.jpg","cid1":4,"cname1":"体育频道","cid2":1,"cname2":"英雄联盟","cid3":24,"cname3":"test","show_status":1,"room_notice":"公告信息","is_vertical":0}]}`
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

	batchRoomInfo := &BatchRoomInfo{
		BaseClient: BaseClient{
			Client: httpClient,
			Secret: "test-secret",
			AID:    "test-aid",
		},
		Token: "test-token",
	}

	timestamp := utils.GetCurrTime()

	if ret, err := batchRoomInfo.do(ts.URL+batchRoomInfoURI, `{"rids": [288016]}`, cast.ToString(timestamp)); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 0 {
			t.Fatal("msg: " + ret.Msg)
		}
		if (ret.Data)[0].RID != 688 {
			t.Fatal("err ret")
		}
	}

	if ret, err := batchRoomInfo.do(ts.URL+batchRoomInfoURI, `{"rids": [288016]}`, "100"); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 100 {
			t.Fatal("msg: " + ret.Msg)
		}
	}
}
