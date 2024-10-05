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

	if rm.CanRequest() {
		rm.requestCounter++
		return nil
	} else {
		return errors.New("rate limit exceeded")
	}
}

func (rm *RateLimitManager) Reset() {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()
	rm.requestCounter = 0
	rm.lastRequest = time.Now()
}

func (rm *RateLimitManager) CanRequest() bool {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	currentTime := time.Now()

	if currentTime.Minute() != rm.lastRequest.Minute() || currentTime.Hour() != rm.lastRequest.Hour() {
		rm.Reset()
	}

	// if requests is under limit, true
	if rm.requestCounter < rm.requestLimit {
		return true
	}

	return false
}
