package maker

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Refactorized to maker.go
type HttpClient struct {
	c *http.Client
}

func NewHttpClient() HttpClient {
	return HttpClient{
		c: &http.Client{Timeout: 10 * time.Second},
	}
}

func (httpClient *HttpClient) CreateRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	httpClient.AddHeaders(req, map[string]string{})
	return req, err
}

func (httpClient *HttpClient) AddHeaders(req *http.Request, headers map[string]string) {
	for keyHeader, keyValue := range headers {
		req.Header.Add(keyHeader, keyValue)
	}
}

func (httpClient *HttpClient) ExecuteRequest(req *http.Request) (*http.Response, error) {
	res, err := httpClient.c.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (httpClient *HttpClient) ReadDataFromBody(body io.Reader) ([]byte, error) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
