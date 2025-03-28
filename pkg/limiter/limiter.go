package limiter

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

// 收到 in 后，将其输出到 out，开启忽略模式，时长 dur
// 忽略模式期间，收到 in，不输出到 out
// 忽略模式结束后，再次收到 in，输出到 out
type limiter struct {
	in     chan any
	out    chan any
	dur    time.Duration
	mu     sync.Mutex // 用于保护 ignore 状态
	ignore bool       // 忽略模式标志
}

func NewLimiter(in chan any, out chan any, dur time.Duration) *limiter {
	if dur <= 0 {
		log.Fatal().Msg("dur must be greater than 0")
	}
	return &limiter{in: in, out: out, dur: dur}
}

func (l *limiter) Start() {
	go func() {
		for msg := range l.in {
			if !l.isIgnore() { // 如果不在忽略模式
				l.out <- msg // 将消息输出到 out
				l.setIgnore(true)
				go func() {
					time.Sleep(l.dur)  // 等待指定的忽略时间
					l.setIgnore(false) // 忽略模式结束
				}()
			}
		}
	}()
}

func (l *limiter) isIgnore() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.ignore
}

func (l *limiter) setIgnore(ignore bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.ignore = ignore
}
