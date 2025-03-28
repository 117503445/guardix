package watcher

import (
	"time"

	"github.com/rs/zerolog/log"
)

type Watcher struct {
	subnet string
}

func NewWatcher(subnet string) *Watcher {
	return &Watcher{subnet: subnet}
}

func (w *Watcher) Start() {
	for {
		arpResults := arp(w.subnet)
		log.Debug().Interface("arpResults", arpResults).Send()

		time.Sleep(time.Minute)
	}
}
