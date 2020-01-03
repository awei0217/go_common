package algorithm

import (
	"fmt"
	"time"
)

/**
漏斗限流
*/
const (
	FUNNEL_INTERVAL = time.Second * 1
)

type FunnelRateLimiter struct {
	interval time.Duration //时间间隔
	cap      int           //漏斗容量
	rate     int           //漏斗流出速率 每秒流多少
	head     int           //放入水的指针
	tail     int           //漏水的指针
	ticker   *time.Ticker  //定时器
}

/**
  cap 漏斗的容量  代表最大容纳多少个请求
  rate 漏斗流出的速度 大于0
  限流速度是 cap - rate
*/
func NewFunnelRateLimiter(cap int, rate int) *FunnelRateLimiter {
	limiter := &FunnelRateLimiter{
		interval: FUNNEL_INTERVAL,
		cap:      cap,
		rate:     rate,
		head:     0,
		tail:     0,
	}
	go leakRate(limiter)
	return limiter
}

/**
  是否允许请求通过（主要看漏斗是否满了）
*/
func (limiter *FunnelRateLimiter) IsAllow() bool {
	if limiter.head-limiter.tail >= limiter.cap { //说明漏斗满了
		return false
	}
	limiter.head++
	return true
}

/**
  模拟漏斗以一定的流速漏水
*/
func leakRate(limiter *FunnelRateLimiter) {
	limiter.ticker = time.NewTicker(limiter.interval)
	for {
		<-limiter.ticker.C

		//根本没有流量，不需要漏
		if limiter.tail >= limiter.head {
			fmt.Println("根本没有流量，不需要漏", limiter.head, limiter.tail)
			continue
		}
		if (limiter.head - limiter.tail) > limiter.rate {
			limiter.tail = limiter.tail + limiter.rate
			fmt.Println("每秒都在流水-1", limiter.head, limiter.tail)
		} else {
			limiter.tail = limiter.head
			fmt.Println("每秒都在流水-2", limiter.head, limiter.tail)
		}
	}
}
