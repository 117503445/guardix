package limiter

import (
	"sync/atomic"
	"time"

	"github.com/rs/zerolog/log"
)

// 收到 in 后，将其输出到 out，开启忽略模式，时长 dur
// 忽略模式期间，收到 in，不输出到 out
// 忽略模式结束后，再次收到 in，输出到 out
type limiter[T any] struct {
	in     chan T
	out    chan T
	dur    time.Duration
	ignore uint32 // atomic flag for ignore mode (0 = false, 1 = true)
}

func NewLimiter[T any](in chan T, out chan T, dur time.Duration) *limiter[T] {
	if dur <= 0 {
		log.Fatal().Msg("dur must be greater than 0")
	}
	return &limiter[T]{in: in, out: out, dur: dur}
}

func (l *limiter[T]) Start() {
	go func() {
		for msg := range l.in {
			if !l.isIgnore() { // 如果不在忽略模式
				l.setIgnore(true)
				go func() {
					time.Sleep(l.dur)  // 等待指定的忽略时间
					l.setIgnore(false) // 忽略模式结束
				}()
				l.out <- msg // 将消息输出到 out
			} else {
				log.Debug().Msg("ignore by limiter")
			}
		}
	}()
}

func (l *limiter[T]) isIgnore() bool {
	return atomic.LoadUint32(&l.ignore) == 1
}

func (l *limiter[T]) setIgnore(ignore bool) {
	var val uint32 = 0
	if ignore {
		val = 1
	}
	atomic.StoreUint32(&l.ignore, val)
}
