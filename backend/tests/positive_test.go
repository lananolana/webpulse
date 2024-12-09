package tests

import (
	"testing"

	"github.com/lananolana/webpulse/backend/tests/testdata"
)

func TestDomainStatus(t *testing.T) {
	for _, url := range testdata.URLs {
		s.CheckSuccessfulDomainStats(t, url)
	}
}
