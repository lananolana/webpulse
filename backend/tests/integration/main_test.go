package tests

import (
	"os"
	"testing"

	"github.com/lananolana/webpulse/backend/tests/integration/client"
	"github.com/lananolana/webpulse/backend/tests/integration/step"
)

var (
	c *client.Client
	s *step.Steps
)

func TestMain(t *testing.M) {

	c = client.NewClient("http://localhost:8080")
	s = step.NewSteps(c)

	run := t.Run()
	os.Exit(run)
}
