package tests

import (
	"os"
	"testing"

	"github.com/lananolana/webpulse/backend/tests/client"
	"github.com/lananolana/webpulse/backend/tests/step"
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
