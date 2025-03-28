package main

import (
	"github.com/117503445/goutils"
	"github.com/117503445/guardix/pkg/cli"
	"github.com/117503445/guardix/pkg/watcher"
	"github.com/rs/zerolog/log"
)

func main() {
	goutils.InitZeroLog()

	cli.CliLoad()

	log.Info().Msg("Starting Guardix")

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

	w := watcher.NewWatcher(cli.Cli.Subnet, pc, phones)
	w.Start()
}
