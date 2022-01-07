package workerpool

import (
	"context"
	"fmt"
	"myhttp/pkg/httpRequester"
	"net/http"
	"testing"
)

func TestNewWorkerPool(t *testing.T) {
	Worker.NewWorkerPool(0)
}

func TestAddTask(t *testing.T) {
	testcases := []struct {
		desc  string
		input []string
	}{
		{
			desc:  "a",
			input: []string{"test"},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			Worker.AddTask(tc.input)
		})
	}
}

type httpMock struct {
	err error
}

func (h httpMock) MakeRequest(ctx context.Context, baseURL string) error {
	return h.err
}

func (h httpMock) NewHttpRequester(client http.Client) {
	return
}

func TestRun(t *testing.T) {
	testHttpRequester := &httpMock{}
	testcases := []struct {
		desc string
		err  error
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

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			testHttpRequester.err = tc.err
			httpRequester.HTTPRequester = testHttpRequester
			Worker.NewWorkerPool(2)
			Worker.AddTask([]string{"test"})

			Worker.Run()
		})
	}

}

