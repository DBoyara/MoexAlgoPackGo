package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DBoyara/MoexAlgoPackGo/pkg/models"
)

type ApiClient struct {
	baseUrl    string
	cert       string
	httpClient *http.Client
}

func NewApiClient(baseUrl, cert string) *ApiClient {
	return &ApiClient{
		baseUrl: baseUrl,
		cert:    cert,
		httpClient: &http.Client{
			Transport: &http.Transport{
				Proxy: nil,
			},
		},
	}
}

func (a *ApiClient) doRequest(url, errText string) (*models.Response, error) {
	resp := &models.Response{}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return resp, err
	}

	request.Header.Add("Cookie", fmt.Sprintf("MicexPassportCert=%s", a.cert))
	request.Header.Add("Cache-Control", "no-cache")

	response, err := a.httpClient.Do(request)
	if err != nil {
		return resp, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("%s | Status Code: %d | Text: %s", errText, response.StatusCode, response.Status)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(responseData, &resp)
	if err != nil {
		return resp, err
	}
	return resp, err
}
