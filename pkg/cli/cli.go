package cli

import (
	"github.com/alecthomas/kong"
	kongtoml "github.com/alecthomas/kong-toml"
	"github.com/rs/zerolog/log"
)

var Cli struct {
	Subnet string `help:"subnet to watch" default:"192.168.1.0/24"`
}

func CliLoad() {
	kong.Parse(&Cli, kong.Configuration(kongtoml.Loader, "config.toml"))
	log.Info().Interface("cli", Cli).Send()
}
