package ws

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type httpClientMock struct {
	mock.Mock
}

func (hClient httpClientMock) Do(r *http.Request) (*http.Response, error) {
	args := hClient.Called()
	return args.Get(0).(*http.Response), args.Error(1)
}
