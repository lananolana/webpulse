package step

import (
	"testing"

	"github.com/lananolana/webpulse/backend/internal/dto"
	"github.com/lananolana/webpulse/backend/tests/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Steps struct {
	*client.Client
}

func NewSteps(client *client.Client) *Steps {
	return &Steps{client}
}

func (s *Steps) CheckSuccessfulDomainStats(t *testing.T, query string) *dto.ServiceStatsResponse {

	response, err := s.Client.GetSiteStatus(t, query)
	require.NoError(t, err, "Failed to fetch stats for %s", query)

	assert.Equal(t, dto.Status(dto.Success), response.Status,
		"Domain: %s failed with unexpected status: %s", query, response.Status)

	return response
}
