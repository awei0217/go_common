package basic

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
)

// 针对web性能监测 ,引入 _ "net/http/pprof" 就可以了
// 浏览器访问:  http://localhost:8080/debug/pprof

//debug/pprof/profile：访问这个链接会自动进行 CPU profiling，持续 30s，并生成一个文件供下载
//debug/pprof/block：Goroutine阻塞事件的记录。默认每发生一次阻塞事件时取样一次。
//debug/pprof/goroutines：活跃Goroutine的信息的记录。仅在获取时取样一次。
//debug/pprof/heap： 堆内存分配情况的记录。默认每分配512K字节时取样一次。
//debug/pprof/mutex: 查看争用互斥锁的持有者。
//debug/pprof/threadcreate: 系统线程创建情况的记录。 仅在获取时取样一次

var Count int64 = 0

// go tool pprof http://127.0.0.1:9909/debug/pprof/profile
func PprofWeb() {
	go func() {
		for {
			time.Sleep(3 * time.Millisecond)
			log.Println(Add("sunpengwei"))
		}
	}()
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}

func handlerData(w http.ResponseWriter, r *http.Request) {
	qUrl := r.URL
	fmt.Println(qUrl)
	fibRev := Fib()
	var fib uint64
	for i := 0; i < 1; i++ {
		fib = fibRev()
		fmt.Println("fib = ", fib)
	}
	str := RandomStr(RandomInt(100, 500))
	str = fmt.Sprintf("Fib = %d; String = %s", fib, str)
	w.Write([]byte(str))
}

func test(w http.ResponseWriter, r *http.Request) {
	fibRev := Fib()
	var fib uint64
	index := Count
	arr := make([]uint64, index)
	var i int64
	for ; i < index; i++ {
		fib = fibRev()
		arr[i] = fib
		fmt.Println("fib = ", fib)
	}
	time.Sleep(time.Millisecond * 500)
	str := fmt.Sprintf("Fib = %v", arr)
	w.Write([]byte(str))
}

func Fib() func() uint64 {
	var x, y uint64 = 0, 1
	return func() uint64 {
		x, y = y, x+y
		return x
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomStr(num int) string {
	seed := time.Now().UnixNano()
	if seed <= 0 {
		seed = time.Now().UnixNano()
	}
	rand.Seed(seed)
	b := make([]rune, num)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func calCount() {
	timeInterval := time.Tick(time.Second)
	for {
		select {
		case i := <-timeInterval:
			Count = int64(i.Second())
		}
	}
}

func PProfCPUApplication() {
	f, _ := os.Create("./cpu.prof")
	pprof.StartCPUProfile(f)

	for i := 1; i < 3000; i++ {
		time.Sleep(3 * time.Millisecond)
		RandomInt(10, 50)
	}
	pprof.StopCPUProfile()
	f.Close()
}

func PProfMemApplication() {
	f, _ := os.Create("./mem.prof")
	for i := 1; i < 3000; i++ {
		time.Sleep(3 * time.Millisecond)
		RandomInt(10, 50)
	}
	pprof.WriteHeapProfile(f)
	f.Close()
}
