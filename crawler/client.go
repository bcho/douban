package crawler

import "net/http"

// HTTPClient performs request.
type HTTPClient struct {
	httpClient *http.Client
}

type httpClientOpts func(*HTTPClient)

// NewHTTPClient creates a HTTPClient instance with given options.
func NewHTTPClient(opts ...httpClientOpts) *HTTPClient {
	c := &HTTPClient{
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c HTTPClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", pickUA())

	return c.httpClient.Do(req)
}

func (c HTTPClient) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

var DefaultClient = NewHTTPClient()
