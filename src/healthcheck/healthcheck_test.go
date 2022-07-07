//go:build unit
// +build unit

package healthcheck

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudoblockio/icon-go-api/api"
	"github.com/sudoblockio/icon-go-api/config"
)

func init() {
	config.ReadEnvironment()
}

func TestHealthCheck(t *testing.T) {
	assert := assert.New(t)

	// Start api
	api.Start()

	// Start healthcheck
	Start()

	resp, err := http.Get("http://localhost:" + config.Config.HealthPort + config.Config.HealthPrefix)
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)
}
