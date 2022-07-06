package http

import (
	"api-google-translate/internal/domain"
	"api-google-translate/internal/maker"
	"api-google-translate/pkg/helper"
	"encoding/json"
	"errors"
	"fmt"
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
	URL_GOOGLE_TRANSLATE           = "https://google-translate1.p.rapidapi.com/language/translate/v2"
	URL_GOOGLE_TRANSLATE_LANGUAGES = "https://google-translate1.p.rapidapi.com/language/translate/v2/languages"
	URL_GOOGLE_TRANSLATE_DETECTOR  = "https://google-translate1.p.rapidapi.com/language/translate/v2/detect"
)

func (httpRepository *httpRepository) GetWordTranslated(word, target, source string) (domain.GoogleTranslate, error) {
	type Translate struct {
		Data struct {
			Translations []domain.GoogleTranslate `json:"translations"`
		} `json:"data"`
	}

	queryString := createQueryStringValues(word, target, source)

	headers, err := getMapWithHeaders(true)
	if err != nil {
		return domain.GoogleTranslate{}, err
	}

	_, responseRaw, err := httpRepository.Rest.MakePostRequest(URL_GOOGLE_TRANSLATE, strings.NewReader(queryString.Encode()), headers)
	if err != nil {
		return domain.GoogleTranslate{}, err
	}

	googleResponse := &Translate{}
	if err = json.Unmarshal(responseRaw, &googleResponse); err != nil {
		return domain.GoogleTranslate{}, err
	}

	return googleResponse.Data.Translations[0], nil
}

func (httpRepository *httpRepository) GetSupportedLanguages() ([]domain.GoogleLanguage, error) {
	type Language struct {
		Data struct {
			Languages []domain.GoogleLanguage `json:"languages"`
		} `json:"data"`
	}

	headers, err := getMapWithHeaders(false)
	if err != nil {
		return []domain.GoogleLanguage{}, err
	}

	_, responseRaw, err := httpRepository.Rest.MakeGetRequest(URL_GOOGLE_TRANSLATE_LANGUAGES, nil, headers)
	if err != nil {
		return []domain.GoogleLanguage{}, err
	}

	googleResponse := &Language{}
	if err = json.Unmarshal(responseRaw, &googleResponse); err != nil {
		return []domain.GoogleLanguage{}, err
	}

	return googleResponse.Data.Languages, nil
}

func (httpRepository *httpRepository) GetDetectedLanguage(text string) (domain.GoogleDetector, error) {
	type Detect struct {
		Data struct {
			Detections [][]struct {
				domain.GoogleDetector
			} `json:"detections"`
		} `json:"data"`
	}

	queryString := make(url.Values)
	queryString.Add("q", text)

	headers, err := getMapWithHeaders(true)
	if err != nil {
		return domain.GoogleDetector{}, err
	}

	_, responseRaw, err := httpRepository.Rest.MakePostRequest(URL_GOOGLE_TRANSLATE_DETECTOR, strings.NewReader(queryString.Encode()), headers)
	if err != nil {
		return domain.GoogleDetector{}, err
	}

	googleResponse := &Detect{}
	if err = json.Unmarshal(responseRaw, &googleResponse); err != nil {
		return domain.GoogleDetector{}, err
	}

	if googleResponse != nil {
		return domain.GoogleDetector{}, errors.New(fmt.Sprintf("Doesn't unmarshal response raw. err: %s", responseRaw))
	}

	return googleResponse.Data.Detections[0][0].GoogleDetector, nil
}

func createQueryStringValues(word string, target string, source string) url.Values {
	queryString := make(url.Values)
	queryString.Add("q", word)
	queryString.Add("target", target)
	queryString.Add("source", source)
	return queryString
}

func getMapWithHeaders(useContentType bool) (map[string]string, error) {
	apiKey, err := helper.GetValueFromEnvironmentVariable("API_GOOGLE_TRANSLATE_KEY")
	if err != nil {
		return nil, err
	}

	apiHost, err := helper.GetValueFromEnvironmentVariable("API_GOOGLE_TRANSLATE_HOST")
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	if useContentType {
		headers["content-type"] = "application/x-www-form-urlencoded"
	}
	headers["Accept-Encoding"] = "application/gzip"
	headers["X-RapidAPI-Key"] = apiKey
	headers["X-RapidAPI-Host"] = apiHost
	return headers, nil
}
