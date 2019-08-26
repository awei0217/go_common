package wechat

import (
	"sync"
	"testing"
)

func TestWeChat(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(1)
	WeChat()
	wg.Wait()
}
