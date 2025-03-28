package cli

import (
	"github.com/alecthomas/kong"
	kongtoml "github.com/alecthomas/kong-toml"
	"github.com/rs/zerolog/log"
)

var Cli struct {
	Subnet  string `help:"subnet to watch" default:"192.168.1.0/24"`
	Devices []*struct {
		Type string `arg:"" help:"device type" enum:"pc,phone"`
		Mac  string `arg:"" help:"mac address"`
		Name string `arg:"" help:"device name"`
	} `help:"devices"`
	Alert struct {
		Endpoint string `help:"alert endpoint"`
		Token    string `help:"alert token"`
	} `embed:"" help:"alert settings" prefix:"alert-"`
}

func CliLoad() {
	kong.Parse(&Cli, kong.Configuration(kongtoml.Loader, "config.toml"))
	log.Info().Interface("cli", Cli).Send()
}
