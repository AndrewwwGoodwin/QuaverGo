package RateLimitManager

import (
	"testing"
	"time"
)

// tests what happens if rate limit is exceeded
func TestRateLimitManagerRateLimitExceeded(t *testing.T) {
	//spin up a full RateLimitManager
	manager := RateLimitManager{requestCounter: 100, requestLimit: 100, lastRequest: time.Now()}
	//try to push another request
	err := manager.Request()
	if err != nil {
		//if the manager throws an error, we're good!
		return
	}
	//otherwise fail the test
	t.Fail()
}

// checks that the Reset() function actually sets the manager to 0 as expected
func TestRateLimitManagerReset(t *testing.T) {
	manager := RateLimitManager{requestCounter: 100, requestLimit: 100, lastRequest: time.Now()}
	manager.Reset()
	if manager.requestCounter == 0 {
		return
	}
	t.Fail()
}

// checks that Request() properly increments the manager
func TestRateLimitManagerRequestIncrement(t *testing.T) {
	manager := RateLimitManager{requestCounter: 85, requestLimit: 100, lastRequest: time.Now()}
	err := manager.Request()
	if err != nil {
		t.FailNow()
	}
	if manager.requestCounter != 86 {
		t.Fail()
	}
	return
}
