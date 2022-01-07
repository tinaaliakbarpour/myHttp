package httpRequester

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"myhttp/config"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)


func TestRequest(t *testing.T) {
	config.Confs = config.Config{
		Debug:  false,
		MyHttp: config.MyHttp{
			TimeOut: 10 * time.Second,
		},
	}
	testcases := []struct {
		desc string
		ctx context.Context
		headers map[string]string
		err error
	}{
		{
			desc:     "a",
			ctx:  context.Background(),
			headers:  nil,
			err:     nil,
		},
		{
			desc:    "b",
			ctx: context.Background(),
			headers: nil,
			err:     fmt.Errorf("faild to do http request"),
		},
	}

	for _,tc := range testcases{
		t.Run(tc.desc, func(t *testing.T) {

			want := "Success!"
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte(want))
			}))
			defer srv.Close()
			baseUrl := srv.URL
			HTTPRequester.NewHttpRequester(http.Client{})

			if tc.desc == "b"{
				config.Confs.MyHttp.TimeOut = 0
			}

			if err := HTTPRequester.MakeRequest(tc.ctx,baseUrl);err != nil{
				assert.Equal(t, err,tc.err)
				return
			}
		})
	}
}