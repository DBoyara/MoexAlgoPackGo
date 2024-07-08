package moexalgopackgo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlgoClient_Authenticate(t *testing.T) {
	tests := []struct {
		name           string
		mockServerFunc func() *httptest.Server
		expectedCert   string
		expectError    bool
	}{
		{
			name: "Successful Authentication",
			mockServerFunc: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					http.SetCookie(w, &http.Cookie{Name: "MicexPassportCert", Value: "test-cert"})
				}))
			},
			expectedCert: "test-cert",
			expectError:  false,
		},
		{
			name: "Failed Authentication Due to HTTP Error",
			mockServerFunc: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "Forbidden", http.StatusForbidden)
				}))
			},
			expectError: true,
		},
		{
			name: "Failed Authentication Due to No Relevant Cookie",
			mockServerFunc: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					// No relevant cookie set
				}))
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := tt.mockServerFunc()
			defer mockServer.Close()

			client := &AlgoClient{
				auth_url: mockServer.URL, // Use mock server URL
				login:    "testuser",
				password: "testpassword",
			}

			err := client.Authenticate()

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCert, client.cert)
			}
		})
	}
}
