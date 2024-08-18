package action

import (
	"github.com/117503445/goutils"
	"github.com/imroc/req/v3"
	"github.com/rs/zerolog/log"

	// "github.com/stretchr/testify/assert"
	"testing"
)

func TestPush(t *testing.T) {
	// ast := assert.New(t)
	req.DevMode() // Treat the package name as a Client, enable development mode

	type Config struct {
		AlertEndpoint string `koanf:"alert_endpoint"`
		AlertToken    string `koanf:"alert_token"`
		AlertChannel  string `koanf:"alert_channel"`
	}
	config := &Config{
		AlertEndpoint: "",
		AlertToken:    "",
		AlertChannel:  "",
	}

	result := goutils.LoadConfig(config)
	result.Dump()

	_, err := req.SetBodyJsonMarshal(
		map[string]string{
			"token":       config.AlertToken,
			"channel":     config.AlertChannel,
			"title":       "GUARDIX",
			"description": "DESCRIPTION",
			"content":     "DING~", // TODO
		}).SetHeader("Content-Type", "application/json").
		Post(config.AlertEndpoint)
	log.Info().Msgf("err: %v", err)
}
