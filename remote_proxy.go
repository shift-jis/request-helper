package request_helper

import (
	"net/url"
	"strings"
)

type RemoteProxy struct {
	Host     string
	Scheme   string
	Username string
	Password string
}

func ParseProxy(proxyUrl string) *RemoteProxy {
	if !strings.Contains(proxyUrl, "://") {
		proxyUrl = "http://" + proxyUrl
	}

	parsedUrl, err := url.Parse(proxyUrl)
	if err != nil {
		return nil
	}

	remoteProxy := &RemoteProxy{Host: parsedUrl.Host, Scheme: parsedUrl.Scheme}
	if parsedUrl.User != nil {
		remoteProxy.Username = parsedUrl.User.Username()
		remoteProxy.Password, _ = parsedUrl.User.Password()
	}
	return remoteProxy
}

func (proxy *RemoteProxy) ShouldAuth() bool {
	return len(proxy.Username) != 0 && len(proxy.Password) != 0
}
