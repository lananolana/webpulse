package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/lananolana/webpulse/backend/internal/config"
	"github.com/lananolana/webpulse/backend/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Client struct {
	BaseURL string
	Client  *http.Client
}

func NewClient(config *config.HTTPServer) *Client {
	baseURL := fmt.Sprintf("http://%s", config.ListenAddr)
	return &Client{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func (c *Client) GetSiteStatus(t *testing.T, queryParam string) *dto.ServiceStatsResponse {

	path := fmt.Sprintf("/api/status?domain=%s", queryParam)
	responseRaw, err := c.get(path)
	require.NoError(t, err, "Failed to get response for domain %s, received body %v", queryParam, responseRaw)
	defer responseRaw.Body.Close()

	assert.Equal(t, http.StatusOK, responseRaw.StatusCode,
		"Expected status code 200 (OK) from our server, but got %d", responseRaw.StatusCode)

	respByte, err := io.ReadAll(responseRaw.Body)
	fmt.Printf("received response body is %s", string(respByte))
	require.NoError(t, err, "Failed to read response body for domain %s", queryParam)

	var response dto.ServiceStatsResponse
	err = json.Unmarshal(respByte, &response)
	require.NoError(t, err, "Failed to unmarshal response body for domain %s", queryParam)

	return &response
}

func (c *Client) get(path string) (*http.Response, error) {
	return c.Client.Get(c.BaseURL + path)
}
