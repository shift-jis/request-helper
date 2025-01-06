package requests

import (
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	request := MustGet("https://tls.peet.ws/api/all")
	request.Header.Set("accept-encoding", "gzip, deflate, br")

	body, _, err := ReadString(http.DefaultClient, request)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(body)
}
