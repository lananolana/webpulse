package tests

import (
	"testing"

	"github.com/lananolana/webpulse/backend/tests/integration/testdata"
)

// The information about provided domains should be returned in full:
// Availability, Performance, Security, Server Info
func TestDomainStatusAndStats(t *testing.T) {
	for _, url := range testdata.SecureURLs {
		s.CheckSuccessfulDomainStats(t, url)
	}
}

// There's no SSL security for the provided domains"
func TestUnsecureStatus(t *testing.T) {
	for _, url := range testdata.UnsecureURLs {
		s.CheckUnsecureDomainStats(t, url)
	}
}

// The error is returned for the domains in wrong format
func TestWrongDomain(t *testing.T) {
	for _, url := range testdata.WrongURLs {
		s.CheckWrongDomainValidation(t, url)
	}
}
