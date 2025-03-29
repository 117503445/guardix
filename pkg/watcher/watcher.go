package watcher

import (
	"fmt"
	"time"

	"github.com/117503445/guardix/pkg/alerter"
	"github.com/117503445/guardix/pkg/common"
	"github.com/rs/zerolog/log"
)

type MAC string

type Metric struct {
	Time  time.Time
	Mac   MAC
	Found bool
	Name  string
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
	alertChan  chan *alerter.Event
}

func NewWatcher(subnet string,
	pc *Device,
	phones []*Device,
	alertChan chan *alerter.Event,
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
		alertChan:  alertChan,
	}
}



func (w *Watcher) Start() {
	for {
		arpResults := arp(w.subnet)
		// log.Debug().Interface("arpResults", arpResults).Send()
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
			if len(w.metricsMap[m.Mac]) > common.METRICS_NUM {
				// keep only last common.METRICS_NUM metrics
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

		counter := make(map[MAC]int)
		for mac, metrics := range w.metricsMap {
			counter[mac] = 0
			for _, metric := range metrics {
				if metric.Found {
					counter[mac]++
				}
			}
		}
		counterName := make(map[string]int)
		for mac := range counter {
			name := ""
			if mac == w.pc.Mac {
				name = w.pc.Name
			} else {
				for _, phone := range w.phones {
					if mac == phone.Mac {
						name = phone.Name
						break
					}
				}
			}
			counterName[name] = counter[mac]
		}
		log.Debug().Interface("counter", counterName).Send()

		// log.Debug().Interface("metricsMap", w.metricsMap).Send()

		if len(w.metricsMap[w.pc.Mac]) == common.METRICS_NUM {
			const THRESHOLD = int(common.METRICS_NUM * 0.95)
			num := 0
			for _, metric := range w.metricsMap[w.pc.Mac] {
				if metric.Found {
					num++
				}
			}
			if num >= THRESHOLD {
				// PC 在 120 次内，有 95% 以上的出现频率

				phoneFound := false
				for _, phone := range w.phones {
					for _, m := range w.metricsMap[phone.Mac] {
						if m.Found {
							phoneFound = true
							break
						}
					}
				}
				if !phoneFound {
					log.Debug().Msg("PC found, but no phone found")
					w.alertChan <- &alerter.Event{
						Message: fmt.Sprintf("[guardix]\nPC %v 长时间在线", w.pc.Name),
					}
				}
			}
		}

		time.Sleep(common.MonitorInterval)
	}
}
