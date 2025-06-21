package cat_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/005_interfaces/001_interfacing_http/001_no_interface/internal/services/cat"
)

func TestGetCatFacts(t *testing.T) {
	tests := map[string]struct {
		url           string
		n             int
		assertionFunc func(t *testing.T, result []string, err error)
	}{
		"n less than 1": {
			n: 0,
			assertionFunc: func(t *testing.T, result []string, err error) {
				assert.Nil(t, result)
				require.Error(t, err)
				assert.EqualError(t, err, "n must be >= 1")
			},
		},
		"invalid url": {
			url: "://invalid-url",
			n:   1,
			assertionFunc: func(t *testing.T, result []string, err error) {
				assert.Nil(t, result)
				require.Error(t, err)
				assert.ErrorContains(t, err, "failed to create request")
			},
		},
		"server unavailable": {
			url: "https://my-invented-url-that-does-not-exist.com/?count=%d",
			n:   1,
			assertionFunc: func(t *testing.T, result []string, err error) {
				assert.Nil(t, result)
				require.Error(t, err)
				assert.ErrorContains(t, err, "failed to make request")
			},
		},
		"successful response": {
			url: "https://meowfacts.herokuapp.com/?count=%d",
			n:   2,
			assertionFunc: func(t *testing.T, result []string, err error) {
				assert.NoError(t, err)
				require.NotNil(t, result)
				assert.Len(t, result, 2)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := &cat.APIHandler{
				Url: tc.url,
			}

			result, err := handler.GetCatFacts(tc.n)

			tc.assertionFunc(t, result, err)
		})
	}
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
