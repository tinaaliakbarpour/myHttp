package httpRequester

import (
	"context"
	"net/http"
)

type Requester interface {
	MakeRequest(ctx context.Context, baseURL string)error
	NewHttpRequester(http.Client)
}

type httpRequester struct{
	client 		http.Client
}



var HTTPRequester Requester = &httpRequester{}
