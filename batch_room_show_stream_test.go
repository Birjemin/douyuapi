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

// TestBatchRoomShowStream
func TestBatchRoomShowStream(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != batchRoomShowStreamURI {
			t.Fatalf("path is invalid: %s, %s'", batchRoomShowStreamURI, path)
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
			raw = `{"code":0,"msg":"","data":{"rid":688,"rtmp_id":0}}`
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

	list := &BatchRoomShowStream{
		BaseClient: BaseClient{
			Client: httpClient,
			Secret: "test-secret",
			AID:    "test-aid",
		},
		Token: "test-token",
	}

	timestamp := utils.GetCurrTime()

	if ret, err := list.do(ts.URL+batchRoomShowStreamURI, `{"rids": [288016]}`, cast.ToString(timestamp)); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 0 {
			t.Fatal("msg: " + ret.Msg)
		}
		if (ret.Data).RID != 688 {
			t.Fatal("err ret")
		}
	}

	if ret, err := list.do(ts.URL+batchRoomShowStreamURI, `{"rids": [288016]}`, "100"); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 100 {
			t.Fatal("msg: " + ret.Msg)
		}
	}
}
