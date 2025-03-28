package watcher

import (
	"encoding/json"
	"strings"

	"github.com/117503445/goutils/gexec"
	"github.com/rs/zerolog/log"
)

type arpResult struct {
	IP     string `json:"ip"`
	Mac    MAC    `json:"mac"`
	Vendor string `json:"vendor"`
}

// subnet: 192.168.1.0/24
func arp(subnet string) []*arpResult {
	results := make([]*arpResult, 0)

	output, err := gexec.Run(
		gexec.Commands(
			[]string{
				"sx", "arp", "--json", subnet,
			},
		),
	)
	if err != nil {
		log.Error().Err(err).Msg("Error getting arp")
	}
	lines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range lines {
		var result *arpResult
		err := json.Unmarshal([]byte(line), &result)
		if err != nil {
			log.Error().Err(err).Str("line", line).Msg("Error unmarshaling arp result")
		} else {
			results = append(results, result)
		}
	}
	return results
}
