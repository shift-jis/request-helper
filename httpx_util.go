package httpx

import (
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
)

func NewGetRequest(url string) *http.Request {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	return request
}

func NewGetRequestWithContext(ctx context.Context, url string) *http.Request {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	return request
}

func NewPostRequest(url string, payload io.Reader) *http.Request {
	request, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		panic(err)
	}
	return request
}

func NewPostRequestWithContext(ctx context.Context, url string, payload io.Reader) *http.Request {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, payload)
	if err != nil {
		panic(err)
	}
	return request
}

func NewPutRequest(url string, payload io.Reader) *http.Request {
	request, err := http.NewRequest(http.MethodPut, url, payload)
	if err != nil {
		panic(err)
	}
	return request
}

func NewPutRequestWithContext(ctx context.Context, url string, payload io.Reader) *http.Request {
	request, err := http.NewRequestWithContext(ctx, http.MethodPut, url, payload)
	if err != nil {
		panic(err)
	}
	return request
}

func NewPatchRequest(url string, payload io.Reader) *http.Request {
	request, err := http.NewRequest(http.MethodPatch, url, payload)
	if err != nil {
		panic(err)
	}
	return request
}

func NewPatchRequestWithContext(ctx context.Context, url string, payload io.Reader) *http.Request {
	request, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, payload)
	if err != nil {
		panic(err)
	}
	return request
}

func NewDeleteRequest(url string, payload io.Reader) *http.Request {
	request, err := http.NewRequest(http.MethodDelete, url, payload)
	if err != nil {
		panic(err)
	}
	return request
}

func NewDeleteRequestWithContext(ctx context.Context, url string, payload io.Reader) *http.Request {
	request, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, payload)
	if err != nil {
		panic(err)
	}
	return request
}

func NewHeadRequest(url string) *http.Request {
	request, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		panic(err)
	}
	return request
}

func NewHeadRequestWithContext(ctx context.Context, url string) *http.Request {
	request, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		panic(err)
	}
	return request
}

func ReadResponseString(client *http.Client, request *http.Request) (string, *http.Response, error) {
	body, response, err := ReadResponseBytes(client, request)
	if err != nil {
		return "", response, err
	}
	return string(body), response, err
}

func ReadResponseBytes(client *http.Client, request *http.Request) ([]byte, *http.Response, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	body, err := ReadResponseBody(response)
	return body, response, err
}

func ReadResponseBody(response *http.Response) ([]byte, error) {
	defer response.Body.Close()

	if encodings, ok := response.Header["Content-Encoding"]; ok && len(encodings) > 0 {
		switch strings.ToLower(encodings[0]) {
		case "br":
			bodyReader := brotli.NewReader(response.Body)
			return io.ReadAll(bodyReader)

		case "gzip":
			bodyReader, err := gzip.NewReader(response.Body)
			if err != nil {
				return nil, err
			}
			defer bodyReader.Close()
			return io.ReadAll(bodyReader)
		}
	}

	return io.ReadAll(response.Body)
}
