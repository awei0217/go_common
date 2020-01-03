package algorithm

import (
	"time"

	"go.uber.org/atomic"
)

//计数器限流 开启一个新的协程监控每经过一秒index清0
type CountLimiter struct {
	count    int64
	unitTime time.Duration
	index    *atomic.Int64
}

//计数器限流，不用新开协程， 每次判断时，
// 先看当前时间和上次时间差是否大于1秒，如果大于则计数器清零0，重新开始,如果小与1秒，则判断计数器到达阈值，返回false，否则返回true
type CountLimiterNew struct {
	count    int64
	lastTime int64
	index    *atomic.Int64
	nano     int64
}

func NewCountLimiter(count int64, unitTime time.Duration) *CountLimiter {
	countLimiter := &CountLimiter{
		count:    count,
		unitTime: unitTime,
		index:    atomic.NewInt64(0),
	}
	go timer(countLimiter)
	return countLimiter
}

func NewCountLimiterNew(count int64, lastTime int64) *CountLimiterNew {
	countLimiterNew := &CountLimiterNew{
		count:    count,
		lastTime: time.Now().UnixNano(),
		index:    atomic.NewInt64(0),
		nano:     1000 * 1000 * 1000,
	}
	return countLimiterNew
}

func (cl *CountLimiterNew) IsAllowNew() bool {
	//已经进入到下一秒中了
	if time.Now().UnixNano()-cl.lastTime > cl.nano {
		cl.lastTime = time.Now().UnixNano()
		cl.index.Store(1)
		return true
	}
	//当前这一秒钟计数器到达阈值了，进行限流
	if cl.index.Load() > cl.count {
		return false
	}
	//计数器加1
	cl.index.Add(1)
	return true
}

func timer(limiter *CountLimiter) {
	ticker := time.NewTicker(limiter.unitTime)
	for {
		<-ticker.C
		limiter.index.Store(0)
	}
}

func (cl *CountLimiter) IsAllow() bool {
	if cl.index.Load() >= cl.count {
		return false
	}
	cl.index.Add(1)
	return true
}
