package cat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// APIHandler is the implementation of the CatAPIHandlerInterface
type APIHandler struct {
	Url        string
	HTTPClient HTTPClient
}

// HTTPClient is a mockable interface for HTTP requests
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// FactsResponse represents the expected response structure from the API
type FactsResponse struct {
	Data []string `json:"data"`
}

func (c *APIHandler) GetCatFacts(n int) ([]string, error) {
	if n < 1 {
		return nil, errors.New("n must be >= 1")
	}

	url := fmt.Sprintf(c.Url, n)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var catFactsResponse FactsResponse
	if err := json.Unmarshal(body, &catFactsResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return catFactsResponse.Data, nil
}
