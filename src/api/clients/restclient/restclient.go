package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enabledMocks = false
	mocks        = make(map[string]*Mock)
)

type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

func getMockId(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

func StartMockups() {
	enabledMocks = true
}

func FlushMockups() {
	mocks = make(map[string]*Mock)
}

func StopMockups() {
	enabledMocks = false
}

func AddMockup(mock Mock) {
	mocks[getMockId(mock.HttpMethod, mock.Url)] = &mock
}

func Post(url string, headers http.Header, body interface{}, ) (*http.Response, error) {
	if enabledMocks {
		mock := mocks[getMockId(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("no mockup found for give request")
		}
		return mock.Response, mock.Err
	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))

	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}

func Get(url string, headers http.Header, queryParam interface{}) (*http.Response, error) {
	if enabledMocks {
		mock := mocks[getMockId(http.MethodGet, url)]
		if mock == nil {
			return nil, errors.New("no mockup found for give request")
		}
		return mock.Response, mock.Err
	}
	//For control over HTTP client headers and other settings, create a Client:
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
