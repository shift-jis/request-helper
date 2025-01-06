package requests

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
)

func MustGet(url string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	return request
}

func MustPost(url string, payload io.Reader) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, url, payload)
	return request
}

func MustPut(url string, payload io.Reader) *http.Request {
	request, _ := http.NewRequest(http.MethodPut, url, payload)
	return request
}

func MustPatch(url string, payload io.Reader) *http.Request {
	request, _ := http.NewRequest(http.MethodPatch, url, payload)
	return request
}

func MustDelete(url string, payload io.Reader) *http.Request {
	request, _ := http.NewRequest(http.MethodDelete, url, payload)
	return request
}

func MustHead(url string) *http.Request {
	request, _ := http.NewRequest(http.MethodHead, url, nil)
	return request
}

func ReadString(client *http.Client, request *http.Request) (string, *http.Response, error) {
	body, response, err := ReadBytes(client, request)
	if err != nil {
		return "", response, err
	}
	return string(body), response, err
}

func ReadBytes(client *http.Client, request *http.Request) ([]byte, *http.Response, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(response.Body)

	body, err := ReadBody(response)
	return body, response, err
}

func ReadBody(response *http.Response) ([]byte, error) {
	if encodings, has := response.Header["Content-Encoding"]; has {
		if strings.EqualFold(encodings[0], "br") {
			reader := brotli.NewReader(response.Body)

			defer func(reader *brotli.Reader, src io.Reader) {
				_ = reader.Reset(src)
			}(reader, response.Body)

			return io.ReadAll(reader)
		} else if strings.EqualFold(encodings[0], "gzip") {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				return nil, err
			}

			defer func(reader *gzip.Reader) {
				_ = reader.Close()
			}(reader)

			return io.ReadAll(reader)
		}
	}
	return io.ReadAll(response.Body)
}
