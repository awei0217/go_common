package algorithm

import (
	"fmt"
	"testing"
	"time"
)

func TestNewCountLimiter(t *testing.T) {

	cl := NewCountLimiter(100, time.Second*1)
	for i := 0; i < 110; i++ {
		fmt.Println(cl.IsAllow())
	}
	time.Sleep(1 * time.Second)
	for i := 0; i < 110; i++ {
		fmt.Println(cl.IsAllow())
	}
}

func TestNewCountLimiterNew(t *testing.T) {
	cl := NewCountLimiterNew(100, time.Now().UnixNano())
	for i := 0; i < 110; i++ {
		fmt.Println(cl.IsAllowNew())
	}
	time.Sleep(1 * time.Second)
	for i := 0; i < 110; i++ {
		fmt.Println(cl.IsAllowNew())
	}
}
