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

	w := watcher.NewWatcher(cli.Cli.Subnet)
	w.Start()
}
