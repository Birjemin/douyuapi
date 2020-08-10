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

// TestGetDanMuList
func TestGetDanMuList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != getDanMuListUri {
			t.Fatalf("path is invalid: %s, %s'", getDanMuListUri, path)
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

		raw := `{"code": 0, "msg":"ok", "data":{"list":[{"room_id":1,"uid":1,"nickname":"","content":"","timestamp":1,"ip":"::0"}],"cnt":1,"page_context":12}}`
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

	list := &GetDanMuList{
		BaseClient: BaseClient{
			Client: httpClient,
			Secret: "test-secret",
			AID:    "test-aid",
		},
		Token: "test-token",
	}

	timestamp := utils.GetCurrTime()
	msg := `{"cid_type":2,"rid":1}`

	if ret, err := list.do(ts.URL+getDanMuListUri, msg, cast.ToString(timestamp)); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 0 || (ret.Data.List)[0].UID != 1 {
			t.Error(errors.New("msg: " + ret.Msg))
		}
	}
}
