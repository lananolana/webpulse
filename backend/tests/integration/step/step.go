package step

import (
	"net/http"
	"testing"

	"github.com/lananolana/webpulse/backend/internal/dto"
	"github.com/lananolana/webpulse/backend/tests/integration/client"
	"github.com/lananolana/webpulse/backend/tests/integration/testdata"
	"github.com/stretchr/testify/assert"
)

type DomainChecks interface {
	CheckSuccessfulDomainStats(t *testing.T, query string)
	CheckUnsecureDomainStats(t *testing.T, query string)
	CheckWrongDomainValidation(t *testing.T, query string)
}

type Steps struct {
	*client.Client
}

func NewSteps(client *client.Client) DomainChecks {
	return &Steps{client}
}

func (s *Steps) CheckSuccessfulDomainStats(t *testing.T, query string) {

	response := s.Client.GetSiteStatus(t, query)

	assert.Equal(t, dto.Status(dto.Success), response.Status,
		"Domain: %s failed with unexpected status: %s", query, response.Status)

	domainShouldBeAvailable(t, response, query)
	performanceShouldBeReturned(t, response, query)
	domainIsSecureInfoShouldBeReturned(t, response, query)
	serverInfoShouldBeReturned(t, response, query)
}

func (s *Steps) CheckUnsecureDomainStats(t *testing.T, query string) {

	response := s.Client.GetSiteStatus(t, query)

	assert.Equal(t, dto.Status(dto.Success), response.Status,
		"Got unexpected status for domain: %s", query)
	domainShouldBeAvailable(t, response, query)
	domainIsNotSecureShouldBeReturned(t, response, query)
}

func (s *Steps) CheckWrongDomainValidation(t *testing.T, query string) {
	response := s.Client.GetSiteStatus(t, query)

	assert.Equal(t, dto.Status(dto.Failed), response.Status,
		"Unexpected status of domain: %s", query)

	assert.Equal(t, testdata.FailMessage, *response.Msg,
		"Unexpected message for domain: %s", testdata.FailMessage, query)
}

func domainShouldBeAvailable(t *testing.T, response *dto.ServiceStatsResponse, query string) {
	assert.NotEmpty(t, response.Availability, "Availability obj is nil for domain: %s", query)

	assert.Equal(t, http.StatusOK, *response.Availability.HTTPStatusCode,
		"Expected 200 status code for domain: %s but got: %d", query, response.Availability.HTTPStatusCode)
}

func performanceShouldBeReturned(t *testing.T, response *dto.ServiceStatsResponse, query string) {
	assert.NotEmpty(t, response.Performance, "Performance is nil for domain: %s", query)

	assert.NotNil(t, response.Performance.ResponseTimeMs, "Response time is nil for domain: %s", query)
	assert.NotNil(t, response.Performance.TransferSpeedKbps, "Transfer speed is nil for domain: %s", query)
	assert.NotNil(t, response.Performance.ResponseSizeKb, "Response size is nil for domain: %s", query)
}

func domainIsSecureInfoShouldBeReturned(t *testing.T, response *dto.ServiceStatsResponse, query string) {
	assert.NotEmpty(t, response.Security, "Security is nil for domain: %s", query)

	assert.NotNil(t, response.Security.SSL, "SSL is nil for domain: %s", query)
	assert.NotNil(t, response.Security.CORS, "CORS is nil for domain: %s", query)
}

func domainIsNotSecureShouldBeReturned(t *testing.T, response *dto.ServiceStatsResponse, query string) {
	assert.NotEmpty(t, response.Security, "Security is nil for domain: %s", query)

	assert.Nil(t, response.Security.SSL, "SSL is not nil when expected for domain: %s", query)
}

func serverInfoShouldBeReturned(t *testing.T, response *dto.ServiceStatsResponse, query string) {
	assert.NotEmpty(t, response.ServerInfo, "Server info is nil for domain: %s", query)

	assert.NotNil(t, response.ServerInfo.IPAddress, "IP address is missing for domain: %s", query)
	assert.NotNil(t, response.ServerInfo.DnsRecords, "DNS records are nil for domain: %s", query)
}
