package limiter

import (
	"sync"
	"time"

	"github.com/ktruedat/recoAssignment/internal/infra/http"
)

type limiter struct {
	mu         sync.Mutex
	tokensChan chan http.Token
}

func (l *limiter) Tries() <-chan http.Token {
	return l.tokensChan
}

const maxReqMinute = 150

func NewLimiter() http.Limiter {
	lm := &limiter{
		mu:         sync.Mutex{},
		tokensChan: make(chan http.Token, maxReqMinute),
	}
	go func() {
		ticker := time.NewTicker(time.Minute)
		for {
			lm.reset()
			<-ticker.C
		}
	}()
	return lm
}

func (l *limiter) reset() {
	l.mu.Lock()
	defer l.mu.Unlock()

	for range maxReqMinute - len(l.tokensChan) {
		l.tokensChan <- http.Token{}
	}
}
