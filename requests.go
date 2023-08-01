package requests

import (
	"compress/gzip"
	"github.com/andybalholm/brotli"
	"io"
	"net/http"
	"strings"
)

func Get(url string) *http.Request {
	request, _ := http.NewRequest("GET", url, nil)
	return request
}

func Post(url string, payload io.Reader) *http.Request {
	request, _ := http.NewRequest("POST", url, payload)
	return request
}

func Put(url string, payload io.Reader) *http.Request {
	request, _ := http.NewRequest("PUT", url, payload)
	return request
}

func Patch(url string, payload io.Reader) *http.Request {
	request, _ := http.NewRequest("PATCH", url, payload)
	return request
}

func Delete(url string, payload io.Reader) *http.Request {
	request, _ := http.NewRequest("DELETE", url, payload)
	return request
}

func DoAndReadString(client *http.Client, request *http.Request) (string, *http.Response, error) {
	body, response, err := DoAndReadByte(client, request)
	if err != nil {
		return "", response, err
	}
	return string(body), response, err
}

func DoAndReadByte(client *http.Client, request *http.Request) ([]byte, *http.Response, error) {
	response, err := Do(client, request)
	if err != nil {
		return nil, nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(response.Body)
	body, err := ReadResponseBody(response)
	return body, response, err
}

func DoAndHandleResponse[T any](client *http.Client, request *http.Request, handler func(response *http.Response, err error) T) T {
	return handler(Do(client, request))
}

func Do(client *http.Client, request *http.Request) (*http.Response, error) {
	return client.Do(request)
}

func SetHeaders(request *http.Request, headerMaps ...map[string]string) {
	for _, headers := range headerMaps {
		for key, value := range headers {
			request.Header.Set(key, value)
		}
	}
}

func ReadResponseBody(response *http.Response) ([]byte, error) {
	if encodings, has := response.Header["content-encoding"]; has {
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
				err = reader.Close()
			}(reader)
			return io.ReadAll(reader)
		}
	}
	return io.ReadAll(response.Body)
}
