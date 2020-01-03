package basic

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//sync 包的学习

//线程安全map学习
/**
添加数据
*/
var m sync.Map

func SyncMapPut(key int, value interface{}) {
	m.Store(key, value)
}

/**
返回键的现有值(如果存在)，否则存储并返回给定的值，如果是读取则返回true，如果是存储返回false。
*/
func SyncMapUpdate(key int, value interface{}) {
	actual, ok := m.LoadOrStore(key, value)
	fmt.Println(actual, ok)

}

/**
删除指定的key
*/
func SyncMapDel(key int) {
	m.Delete(key)
}

/**
读取key的值，如果不存在，返回 nil,false
*/
func SyncMapRead(key int) {
	v, ok := m.Load(key)
	fmt.Println(v, ok)
}
func SyncMapRange() {
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

//互斥锁
var mutex sync.Mutex

func StudyMutex() {

	wait := sync.WaitGroup{}
	wait.Add(1)
	mutex.Lock()

	fmt.Println("等待.....")

	go func() {
		mutex.Lock()
	}()
	time.Sleep(1 * time.Second)
	mutex.Unlock()

}

func StudySyncOnce() {

	val := 0
	once := sync.Once{}
	f := func() {
		val++
	}
	for i := 0; i < 10; i++ {
		//虽然循环了10次，但只是执行了一次
		once.Do(f)
	}
	fmt.Println(val) //结果是1
}

func StudySyncPool() {
	runtime.GOMAXPROCS(0)
	pool := sync.Pool{}

	pool.Put(1)

	delta := 1
	t := uint64(delta) << 32
	s := uint64(0)
	state := atomic.AddUint64(&s, t)
	t1 := int32(state << 32)
	w := uint32(state)
	fmt.Println(t, t1, w, state)
}
