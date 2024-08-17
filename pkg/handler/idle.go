package trigger

import (
	"time"

	"github.com/lextoumbourou/idle"
	"github.com/rs/zerolog/log"
)

type IdleHandler struct {
	durationThreshhold time.Duration
}

func NewIdleHandler() *IdleHandler {
	return &IdleHandler{
		durationThreshhold: 1 * time.Hour,
	}
}

// Duration returns the time since the last user input
func (i *IdleHandler) Passed() bool {
	d, err := i.lastInputDuration()
	if err != nil {
		log.Error().Err(err).Msg("Error getting idle time")
		return false
	}
	return d > i.durationThreshhold
}

// Duration returns the time since the last user input
func (i *IdleHandler) lastInputDuration() (time.Duration, error) {
	return idle.Get()
}
