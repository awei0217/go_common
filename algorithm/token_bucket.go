package algorithm

import (
	"fmt"
	"sync"
	"time"
)

/**
令牌桶限流算法
*/
const (
	/**
	默认时间间隔
	*/
	INTERVAL = time.Second * 1
	/**
	时间间隔内放入令牌的个数(比如时间间隔为1秒，那就是每隔一秒放入100 个令牌)
	*/
	INTERVAL_IN_TOKEN = 100
	/**
	令牌桶初始时的令牌个数
	*/
	AVAIL_TOKEN = 100
)

type TokenBucket struct {
	interval        time.Duration //时间间隔
	ticker          *time.Ticker  //定时器
	cap             int           // 桶容量
	avail           int           //桶内一开始令牌数
	tokenArray      []int         //存储令牌的数组
	intervalInToken int           //时间间隔内放入令牌的个数
	index           int           //数组放入令牌的下标处
	mutex           sync.Mutex
}

/**
创建一个令牌通
入参为令牌桶的容量
*/
func NewTokenBucket(cap int) *TokenBucket {

	tokenBucket := &TokenBucket{
		interval:        INTERVAL,
		cap:             cap,
		avail:           AVAIL_TOKEN,
		tokenArray:      make([]int, cap, cap),
		intervalInToken: INTERVAL_IN_TOKEN,
		index:           0,
		mutex:           sync.Mutex{},
	}
	go adjustTokenDaemon(tokenBucket)
	return tokenBucket
}

/**
调整令牌桶令牌的协程
*/
func adjustTokenDaemon(tokenBucket *TokenBucket) {
	//如果桶内一开始的令牌小于初始令牌，开始放入初始令牌
	for tokenBucket.index < tokenBucket.avail {
		tokenBucket.tokenArray[tokenBucket.index] = 1
		tokenBucket.index++
	}
	tokenBucket.ticker = time.NewTicker(INTERVAL)
	go func(t *time.Ticker) {
		for {
			<-t.C
			putToken(tokenBucket)
		}
	}(tokenBucket.ticker)
}
func putToken(tokenBucket *TokenBucket) {
	tokenBucket.mutex.Lock()
	for i := 0; i < tokenBucket.intervalInToken; i++ {
		if tokenBucket.index > tokenBucket.cap-1 {
			fmt.Println("令牌通已满")
			break
		}
		tokenBucket.tokenArray[tokenBucket.index] = 1
		tokenBucket.index++
		fmt.Println("放入令牌成功")
	}
	defer tokenBucket.mutex.Unlock()
}

/**
从令牌桶弹出一个令牌，如果令牌通有令牌，返回true，否则返回false
*/
func (tokenBucket *TokenBucket) PopToken() bool {
	defer tokenBucket.mutex.Unlock()
	tokenBucket.mutex.Lock()
	if tokenBucket.index <= 0 {
		return false
	}
	tokenBucket.tokenArray[tokenBucket.index-1] = 0
	tokenBucket.index--
	return true
}
