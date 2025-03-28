package watcher

import (
	"time"

	"github.com/rs/zerolog/log"
)

type MAC string

type Metric struct {
	Time  time.Time
	Mac   MAC
	Found bool
	Name string
}

type Device struct {
	Mac  MAC
	Name string
}

type Watcher struct {
	subnet string
	pc     *Device
	phones []*Device

	metricsMap map[MAC][]*Metric
}

func NewWatcher(subnet string,
	pc *Device,
	phones []*Device,
) *Watcher {
	log.Debug().Interface("subnet", subnet).Interface("pc", pc).Interface("phones", phones).
		Msg("NewWatcher")

	metricsMap := make(map[MAC][]*Metric)
	metricsMap[pc.Mac] = make([]*Metric, 0)
	for _, phone := range phones {
		metricsMap[phone.Mac] = make([]*Metric, 0)
	}

	return &Watcher{subnet: subnet,
		pc:         pc,
		phones:     phones,
		metricsMap: metricsMap,
	}
}

func (w *Watcher) Start() {
	for {
		arpResults := arp(w.subnet)
		log.Debug().Interface("arpResults", arpResults).Send()
		curMacs := make(map[MAC]bool)
		for _, arpResult := range arpResults {
			curMacs[arpResult.Mac] = true
		}

		found := func(mac MAC) bool {
			ok := curMacs[mac]
			return ok
		}

		appendMetric := func(m *Metric) {
			w.metricsMap[m.Mac] = append(w.metricsMap[m.Mac], m)
			if len(w.metricsMap[m.Mac]) > 2*60 {
				// keep only last 2 hours of metrics
				w.metricsMap[m.Mac] = w.metricsMap[m.Mac][1:]
			}
		}
		appendMetric(
			&Metric{
				Time:  time.Now(),
				Mac:   w.pc.Mac,
				Found: found(w.pc.Mac),
				Name:  w.pc.Name,
			},
		)
		for _, phone := range w.phones {
			appendMetric(
				&Metric{
					Time:  time.Now(),
					Mac:   phone.Mac,
					Found: found(phone.Mac),
					Name:  phone.Name,
				},
			)
		}

		log.Debug().Interface("metricsMap", w.metricsMap).Send()

		time.Sleep(time.Minute)
	}
}
