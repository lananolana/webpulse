package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/lananolana/webpulse/backend/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Client struct {
	Host   string
	Client *http.Client
}

func NewClient(host string) *Client {
	return &Client{
		Host:   host,
		Client: &http.Client{},
	}
}

func (c *Client) GetSiteStatus(t *testing.T, queryParam string) (*dto.ServiceStatsResponse, error) {

	path := fmt.Sprintf("/status?domain=%s", queryParam)
	responseRaw, err := c.get(path)
	require.NoError(t, err, "Failed to get response")
	defer responseRaw.Body.Close()

	assert.Equal(t, http.StatusOK, responseRaw.StatusCode,
		"Expected status code 200 (OK) from our server, but got %d", responseRaw.StatusCode)

	respByte, err := io.ReadAll(responseRaw.Body)
	require.NoError(t, err, "Failed to read response body")

	var response dto.ServiceStatsResponse
	err = json.Unmarshal(respByte, &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	return &response, nil
}

func (c *Client) get(path string) (*http.Response, error) {

	return c.Client.Get(c.Host + path)
}
