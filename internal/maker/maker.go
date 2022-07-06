package maker

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Rest interface {
	MakeGetRequest(url string, body io.Reader, headers map[string]string) (int, []byte, error)
	MakePostRequest(url string, body io.Reader, headers map[string]string) (int, []byte, error)
	MakeDeleteRequest(url string, body io.Reader, headers map[string]string) (int, []byte, error)
	MakePutRequest(url string, body io.Reader, headers map[string]string) (int, []byte, error)
}

type rest struct {
	c *http.Client
}

func NewRest() Rest {
	return &rest{
		&http.Client{Timeout: 10 * time.Second},
	}
}

func (r *rest) MakeGetRequest(url string, body io.Reader, headers map[string]string) (int, []byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, body)
	if err != nil {
		return 0, nil, err
	}

	r.AddHeaders(req, headers)

	res, err := r.c.Do(req)
	if err != nil {
		return 0, nil, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, data, nil
}

func (r *rest) MakePostRequest(url string, body io.Reader, headers map[string]string) (int, []byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return 0, nil, err
	}

	r.AddHeaders(req, headers)

	res, err := r.c.Do(req)
	if err != nil {
		return 0, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, data, nil
}

//TODO: Implements methods
func (r *rest) MakeDeleteRequest(url string, body io.Reader, headers map[string]string) (int, []byte, error) {
	return 0, nil, nil
}

func (r *rest) MakePutRequest(url string, body io.Reader, headers map[string]string) (int, []byte, error) {
	return 0, nil, nil
}

func (r *rest) AddHeaders(req *http.Request, headers map[string]string) {
	for keyHeader, keyValue := range headers {
		req.Header.Add(keyHeader, keyValue)
	}
}
