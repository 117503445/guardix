package action

import (
	"github.com/imroc/req/v3"
)

type Config struct {
	AlertEndpoint string
	AlertToken    string
	AlertChannel  string
}

func Push(config Config) error {
	_, err := req.SetBodyJsonMarshal(
		map[string]string{
			"token":       config.AlertToken,
			"channel":     config.AlertChannel,
			"title":       "TITLE",
			"description": "DESCRIPTION",
			"content":     "DING~", // TODO
		}).SetHeader("Content-Type", "application/json").
		Post(config.AlertEndpoint)
	return err
}
