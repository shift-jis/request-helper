package requests

import (
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	request := Get("https://www.google.com")
	SetHeaders(request, map[string]string{
		"accept-encoding": "gzip, deflate, br",
	})
	if err := DoAndHandleResponse[error](http.DefaultClient, request, func(response *http.Response, err error) error {
		body, err := ReadResponseBody(response)
		println(string(body))
		return err
	}); err != nil {
		panic(err)
	}
}
