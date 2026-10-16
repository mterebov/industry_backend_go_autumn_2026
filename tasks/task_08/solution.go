package main

import (
	"sync"
	"time"
)

type Clock interface{ Now() time.Time }
type Limiter struct {
	mu     sync.Mutex
	clock  Clock
	rate   float64
	burst  int
	tokens float64
	last   time.Time
}

func NewLimiter(clock Clock, ratePerSec float64, burst int) *Limiter {
	var limiter Limiter
	limiter.clock = clock
	limiter.rate = ratePerSec
	limiter.burst = burst
	limiter.tokens = float64(burst)
	return &limiter
}
func (l *Limiter) AllowN(n int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.burst <= 0 || l.clock == nil || n <= 0 || n > l.burst {
		return false
	}

	now := l.clock.Now()
	if now.After(l.last) { // часы назад -> ничего не делаем, last не трогаем
		if l.rate > 0 {
			l.tokens += now.Sub(l.last).Seconds() * l.rate
			if l.tokens > float64(l.burst) {
				l.tokens = float64(l.burst)
			}
		}
		l.last = now
	}

	if l.tokens >= float64(n) {
		l.tokens -= float64(n)
		return true
	}
	return false
}
