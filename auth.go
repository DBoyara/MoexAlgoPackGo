package moexalgopackgo

import (
	"fmt"
	"net/http"
)

func (a *AlgoClient) Authenticate() error {
	request, err := http.NewRequest(http.MethodGet, a.auth_url, nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(a.login, a.password)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to authenticate | Status Code: %d | Text: %s", response.StatusCode, response.Status)
	}

	cookies := response.Cookies()

	for _, cookie := range cookies {
		if cookie.Name == "MicexPassportCert" {
			a.cert = cookie.Value
			break
		}
	}

	return nil
}
