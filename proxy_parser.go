package requests

import (
	"net/url"
	"strings"
)

func ParseProxy(proxyUrl string) (*url.URL, error) {
	if !strings.Contains(proxyUrl, "://") {
		proxyUrl = "http://" + proxyUrl
	}
	return url.Parse(proxyUrl)
}
