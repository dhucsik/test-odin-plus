package parser

import "net/http"

type customUserAgentTransport struct {
	transport http.RoundTripper
	userAgent string
}

func (t *customUserAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.userAgent)
	return t.transport.RoundTrip(req)
}

func GetHttpClient() *http.Client {
	transport := &http.Transport{}

	transport.RegisterProtocol("http", &customUserAgentTransport{
		transport: http.DefaultTransport,
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	})

	client := &http.Client{
		Transport: transport,
	}

	return client
}
