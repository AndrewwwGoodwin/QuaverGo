package RateLimitManager

import (
	"errors"
	"sync"
	"time"
)

type RateLimitManager struct {
	requestCounter int
	requestLimit   int
	lastRequest    time.Time
	mutex          sync.Mutex
}

func Init(RPM int) *RateLimitManager {
	return &RateLimitManager{requestCounter: 0, requestLimit: RPM, lastRequest: time.Now()}
}

func (rm *RateLimitManager) Request() error {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	// Reset if a minute has passed
	if time.Since(rm.lastRequest) >= time.Minute {
		rm.Reset()
	}

	if rm.requestCounter < rm.requestLimit {
		rm.requestCounter++
		return nil
	}
	return errors.New("rate limit exceeded")
}

func (rm *RateLimitManager) Reset() {
	rm.requestCounter = 0
	rm.lastRequest = time.Now()
}
