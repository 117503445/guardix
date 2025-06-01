package alerter

import (
	"github.com/imroc/req/v3"
	"github.com/rs/zerolog/log"
)

type Alerter struct {
	endpoint string
	token    string
}

type Event struct {
	Message string
}

func NewAlerter(endpoint string, token string) *Alerter {
	a := &Alerter{endpoint: endpoint,
		token: token,
	}

	return a
}

func (a *Alerter) Alert(message string) {
	log.Info().Str("message", message).Msg("Alerting")

	_, err := req.SetBodyJsonMarshal(
		map[string]string{
			"token":       a.token,
			"title":       "[guardix] 警告",
			"description": "DESCRIPTION",
			"content":     message,
		}).SetHeader("Content-Type", "application/json").
		Post(a.endpoint)
	if err != nil {
		log.Error().Err(err).Msg("alert error")
	}
}

func (a *Alerter) TestAlert() {
	log.Info().Msg("Testing alert")

	testClient := req.C()
	testClient.DevMode()

	_, err := testClient.R().SetBodyJsonMarshal(
		map[string]string{
			"token":       a.token,
			"title":       "TITLE",
			"description": "DESCRIPTION",
			"content":     "[guardix]\n Alert 测试",
		}).SetHeader("Content-Type", "application/json").
		Post(a.endpoint)
	if err != nil {
		log.Fatal().Err(err).Msg("alert error")
	}
}
