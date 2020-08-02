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

// TestToken
func TestGetRoomInfo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != RoomInfoUri {
			t.Fatalf("path is invalid: %s, %s'", RoomInfoUri, path)
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

		raw := `{"code":0,"msg":"","data":{"rid":688,"room_src":"room_src","room_src_max":"room_src_max","room_name":"房间","hn":6030,"nickname":"test","avatar":"a7e997_big.jpg","cid1":4,"cname1":"体育频道","cid2":1,"cname2":"英雄联盟","cid3":24,"cname3":"test","show_status":1,"show_time":1596358821,"unuid":"","room_notice":"公告信息","is_vertical":0,"fans":0}}`
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

	room := &RoomInfo{
		BaseClient: BaseClient{
			Client: httpClient,
			Secret: "test-secret",
			AID:    "test-aid",
		},
		Token: "test-token",
	}

	timestamp := utils.GetCurrTime()

	if ret, err := room.do(ts.URL+RoomInfoUri, `{"rid": 288016}`, cast.ToString(timestamp)); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 0 {
			t.Fatal("msg: " + ret.Msg)
		}
		if ret.Data.RID != 688 {
			t.Fatal("err ret")
		}
	}
}
