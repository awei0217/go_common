package algorithm

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewTokenBucket(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	tokenBucket := NewTokenBucket(1000)
	for {
		//模拟每隔10毫秒消费一个令牌
		fmt.Println(tokenBucket.PopToken())
		time.Sleep(time.Millisecond * 10)
	}
	wg.Wait()
}


