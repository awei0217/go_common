package basic

import (
	"fmt"
	"sync/atomic"
)

func AtomicStudy() {
	var m int32 = 9

	// 给i 加 2 返回加以后的值
	new := atomic.AddInt32(&m, 2)
	fmt.Println(new)

}
