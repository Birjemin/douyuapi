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

// TestGetAudioPlay
func TestGetAudioPlay(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != getAudioPlayUri {
			t.Fatalf("path is invalid: %s, %s'", getAudioPlayUri, path)
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

		raw := `{"code":0,"msg":"","data":[{"rid":1,"uid":1,"rtmpCdn":"","rtmpLive":"","rtmpURL":1}]}`
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

	allLive := &AllLive{
		BaseClient: BaseClient{
			Client: httpClient,
			Secret: "test-secret",
			AID:    "test-aid",
		},
		Token: "test-token",
	}

	timestamp := utils.GetCurrTime()

	msg := `{"rid":1}`

	if ret, err := allLive.do(ts.URL+getAudioPlayUri, msg, cast.ToString(timestamp)); err != nil {
		t.Error(err)
	} else {
		if ret.Code != 0 || len(ret.Data) != 1 {
			t.Error(errors.New("msg: " + ret.Msg))
		}
	}
}