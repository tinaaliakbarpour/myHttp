package main

import (
	"context"
	"fmt"
	"myhttp/pkg/httpRequester"
	"myhttp/pkg/workerpool"
	"net/http"
	"testing"
)

//mock httpRequester package
type httpRequesterMock struct {
	err error
}

func (h httpRequesterMock) MakeRequest(ctx context.Context, baseURL string) error {
	return h.err
}

func (h httpRequesterMock) NewHttpRequester(client http.Client) {
	return
}

//mock workerpool package
type workerMock struct {

}

func (w workerMock) NewWorkerPool(maxWorker int) {
	return
}

func (w workerMock) Run() {
	return
}

func (w workerMock) AddTask(task []string) {
	return
}


//test main function
func Test(t *testing.T) {
	testHttp := &httpRequesterMock{}
	testWorker := &workerMock{}
	testcases := []struct{
		desc string
		err error
	}{
		{
			desc: "a",
			err:  nil,
		},
		{
			desc: "b",
			err:  fmt.Errorf("error"),
		},
	}

	for _ , tc := range testcases{
		t.Run(tc.desc , func(t *testing.T) {
			testHttp.err = tc.err
			httpRequester.HTTPRequester = testHttp
			workerpool.Worker = testWorker
			main()
		})
	}
}