package tests

import (
	"os"
	"testing"

	"github.com/lananolana/webpulse/backend/internal/config"
	"github.com/lananolana/webpulse/backend/tests/integration/client"
	"github.com/lananolana/webpulse/backend/tests/integration/step"
)

const (
	localhost  = "http://localhost:8080"
	configPath = "../../configs/app.yaml"
)

var (
	c *client.Client
	s step.DomainChecks
)

func TestMain(t *testing.M) {

	config := config.MustLoad(configPath)

	c = client.NewClient(&config.App.HTTPServer)
	s = step.NewSteps(c)

	run := t.Run()
	os.Exit(run)
}
