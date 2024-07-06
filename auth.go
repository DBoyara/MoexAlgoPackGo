package moexalgopackgo

import (
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

	cookies := response.Cookies()

	for _, cookie := range cookies {
		if cookie.Name == "MicexPassportCert" {
			a.cert = cookie.Value
			break
		}
	}

	return nil
}
