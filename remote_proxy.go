package request_helper

type RemoteProxy struct {
	Host     string
	Scheme   string
	Username string
	Password string
}

func (proxy *RemoteProxy) ShouldAuth() bool {
	return len(proxy.Username) != 0 && len(proxy.Password) != 0
}
