package http

import (
	"api-google-translate/internal/maker"
	"api-google-translate/pkg/helper"
	"net/url"
	"strings"
)

type httpRepository struct {
	Rest maker.Rest
}

func NewHttpRepository() interface{} {
	return &httpRepository{
		Rest: maker.NewRest(),
	}
}

const (
	URL_MOTIVATIONS = "https://motivational-quotes1.p.rapidapi.com/motivation"
)

func (httpRepository *httpRepository) GetMotivation() (string, error) {
	body := generateBody()
	headers, err := getMapWithHeaders()
	if err != nil {
		return "", err
	}

	_, responseRaw, err := httpRepository.Rest.MakePostRequest(URL_MOTIVATIONS, strings.NewReader(body.Encode()), headers)
	if err != nil {
		return "", err
	}

	return string(responseRaw), nil
}

func generateBody() url.Values {
	body := make(url.Values)
	body.Add("key1", "value")
	body.Add("key2", "value")
	return body
}

func getMapWithHeaders() (map[string]string, error) {
	apiKey, err := helper.GetValueFromEnvironmentVariable("API_MOTIVATIONS_KEY")
	if err != nil {
		return nil, err
	}

	apiHost, err := helper.GetValueFromEnvironmentVariable("API_MOTIVATIONS_HOST")
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["content-type"] = "application/json"
	headers["X-RapidAPI-Key"] = apiKey
	headers["X-RapidAPI-Host"] = apiHost
	return headers, nil
}
