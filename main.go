package main

import (
	"time"

	"github.com/117503445/goutils"
	"github.com/117503445/guardix/pkg/action"
	"github.com/117503445/guardix/pkg/handler"
	"github.com/rs/zerolog/log"
)

func main() {
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

	idleHandler := handler.NewIdleHandler()
	netHandler := handler.NewNetHandler(4)
	nets := []bool{}

	var lastAlertTime time.Time

	for {
		nets = append(nets, netHandler.Passed())
		if len(nets) > 5 {
			nets = nets[1:]
		}
		isNetOk := true
		for _, net := range nets {
			if !net {
				isNetOk = false
				break
			}
		}

		isIdleOk := idleHandler.Passed()

		log.Info().Bool("isNetOk", isNetOk).Bool("isIdleOk", isIdleOk).Msg("Check")
		if isNetOk && isIdleOk {
			if time.Since(lastAlertTime) > time.Hour {
				lastAlertTime = time.Now()
				action.Push(action.Config{
					AlertEndpoint: config.AlertEndpoint,
					AlertToken:    config.AlertToken,
					AlertChannel:  config.AlertChannel,
				})
			}
		}

		time.Sleep(time.Minute)
	}
}
