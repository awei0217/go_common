package algorithm

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewFunnelRateLimiter(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	limiter := NewFunnelRateLimiter(1000, 1000)
	for {
		fmt.Println(limiter.IsAllow())
		time.Sleep(2 * time.Millisecond)
	}
	wg.Wait()
}
