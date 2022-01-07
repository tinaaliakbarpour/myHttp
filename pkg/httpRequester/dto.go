package httpRequester

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"myhttp/config"
	"net/http"
)

// NewHttpRequester will fill the logger and client fields
func (h *httpRequester) NewHttpRequester(client http.Client) {
	h.client = client
}

//MakeRequest will make http request concurrently
func (h *httpRequester) MakeRequest(ctx context.Context, baseURL string) error {
	// create deadline context
	ctxWithTime, cancel := context.WithTimeout(ctx, config.Confs.MyHttp.TimeOut)
	defer cancel()
	request, err := http.NewRequestWithContext(ctxWithTime, "GET", baseURL, nil)
	if err != nil {
		log.Printf("url %s [MakeRequest]failed with error: %s",baseURL,err.Error())
		return fmt.Errorf("failed to make a new http request")
	}
	resp, err := h.client.Do(request)
	if err != nil {
		log.Printf("url %s [MakeRequest]failed with error: %s",baseURL,err.Error())
		return fmt.Errorf("faild to do http request")
	}
	defer resp.Body.Close()

	byteCode, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("url %s [MakeRequest]failed with error: %s",baseURL,err.Error())
		return fmt.Errorf("failed in fuction ioutil.ReadAll")
	}
	fmt.Printf("%s %x \n", baseURL, md5.Sum(byteCode))
	return nil
}
