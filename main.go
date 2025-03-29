package main

import (
	"github.com/117503445/goutils"
	"github.com/117503445/guardix/pkg/alerter"
	"github.com/117503445/guardix/pkg/cli"
	"github.com/117503445/guardix/pkg/common"
	"github.com/117503445/guardix/pkg/limiter"
	"github.com/117503445/guardix/pkg/watcher"
	"github.com/rs/zerolog/log"
)

func main() {
	goutils.InitZeroLog(goutils.WithProduction{})

	cli.CliLoad()
	log.Info().Msg("Starting Guardix")
	a := alerter.NewAlerter(cli.Cli.Alert.Endpoint, cli.Cli.Alert.Token)
	a.TestAlert()

	var pc *watcher.Device
	phones := make([]*watcher.Device, 0)
	for _, device := range cli.Cli.Devices {
		if device.Type == "pc" {
			pc = &watcher.Device{
				Mac:  watcher.MAC(device.Mac),
				Name: device.Name,
			}
		} else {
			phones = append(phones, &watcher.Device{
				Mac:  watcher.MAC(device.Mac),
				Name: device.Name,
			})
		}
	}

	alertChan := make(chan *alerter.Event)
	LimitedChan := make(chan *alerter.Event)

	w := watcher.NewWatcher(cli.Cli.Subnet, pc, phones, alertChan)
	go w.Start()

	l := limiter.NewLimiter(alertChan, LimitedChan, common.AlertMuteDuration)
	go l.Start()

	for alertEvent := range LimitedChan {
		a.Alert(alertEvent.Message)
	}
}
