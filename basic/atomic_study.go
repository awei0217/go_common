package basic

import (
	"fmt"
	"sync/atomic"
)

func AtomicStudy() {
	var m int32 = 9
	var i *int32 = &m
	// 给i 加 2 返回加以后的值
	new := atomic.AddInt32(i, 2)
	fmt.Println(new)


}
