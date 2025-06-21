package cat_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/005_interfaces/001_interfacing_http/002_with_interface/internal/services/cat"
)

// MockHTTPClient is a mock implementation of the HTTPClient interface
type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do calls the configured DoFunc
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return nil, nil
}

func TestGetCatFacts(t *testing.T) {
	tests := map[string]struct {
		url           string
		httpClient    cat.HTTPClient
		n             int
		assertionFunc func(t *testing.T, result []string, err error)
	}{
		"n less than 1": {
			n:          0,
			httpClient: http.DefaultClient,
			assertionFunc: func(t *testing.T, result []string, err error) {
				assert.Nil(t, result)
				require.Error(t, err)
				assert.EqualError(t, err, "n must be >= 1")
			},
		},
		"invalid url": {
			url:        "://invalid-url",
			n:          1,
			httpClient: http.DefaultClient,
			assertionFunc: func(t *testing.T, result []string, err error) {
				assert.Nil(t, result)
				require.Error(t, err)
				assert.ErrorContains(t, err, "failed to create request")
			},
		},
		"server unavailable": {
			url:        "https://my-invented-url-that-does-not-exist.com/?count=%d",
			n:          1,
			httpClient: http.DefaultClient,
			assertionFunc: func(t *testing.T, result []string, err error) {
				assert.Nil(t, result)
				require.Error(t, err)
				assert.ErrorContains(t, err, "failed to make request")
			},
		},
		"successful response": {
			url: "https://meowfacts.herokuapp.com/?count=%d",
			n:   2,
			httpClient: &MockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					responseBody := `{"data": ["Cats are great!", "Cats love to sleep."]}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(responseBody)),
					}, nil
				},
			},
			assertionFunc: func(t *testing.T, result []string, err error) {
				assert.NoError(t, err)
				require.NotNil(t, result)
				assert.Len(t, result, 2)
				assert.Equal(t, result[0], "Cats are great!")
			},
		},
		"unexpected status code": {
			url: "https://meowfacts.herokuapp.com/?count=%d",
			n:   1,
			httpClient: &MockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusInternalServerError,
						Body:       io.NopCloser(strings.NewReader("Internal Server Error")),
					}, nil
				},
			},
			assertionFunc: func(t *testing.T, result []string, err error) {
				assert.Nil(t, result)
				require.Error(t, err)
				assert.ErrorContains(t, err, "unexpected status code: 500")
			},
		},
		"invalid response body": {
			url: "https://meowfacts.herokuapp.com/?count=%d",
			n:   1,
			httpClient: &MockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader("Invalid JSON")),
					}, nil
				},
			},
			assertionFunc: func(t *testing.T, result []string, err error) {
				assert.Nil(t, result)
				require.Error(t, err)
				assert.ErrorContains(t, err, "failed to unmarshal response")
			},
		},
		"body can't be read": {
			url: "https://meowfacts.herokuapp.com/?count=%d",
			n:   1,
			httpClient: &MockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(&errorReader{}), // Custom reader that always returns an error
					}, nil
				},
			},
			assertionFunc: func(t *testing.T, result []string, err error) {
				assert.Nil(t, result)
				require.Error(t, err)
				assert.ErrorContains(t, err, "failed to read response body")
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := &cat.APIHandler{
				HTTPClient: tc.httpClient,
				Url:        tc.url,
			}

			result, err := handler.GetCatFacts(tc.n)

			tc.assertionFunc(t, result, err)
		})
	}
}

type errorReader struct{}

// Read implements the io.Reader interface and always returns an error
func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, io.ErrUnexpectedEOF // Always return an error
}

// equalSlices checks if two slices of strings are equal
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
