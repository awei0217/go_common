package benchmark

import (
	"fmt"
	"time"
)

func Fib(n int64) int64 {
	if n <= 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

type I interface {
	f()
}

type S struct {
	a int64
}

var a int

func f() {
	a++
}

func Method() {
	start := time.Now().UnixNano()
	for j := 0; j < 1; j++ {
		f()
	}
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

func bubbleSort(arr *[]int) {
	for j := 0; j < len(*arr)-1; j++ {
		for k := 0; k < len(*arr)-1-j; k++ {
			if (*arr)[k] < (*arr)[k+1] {
				temp := (*arr)[k]
				(*arr)[k] = (*arr)[k+1]
				(*arr)[k+1] = temp
			}
		}
	}
}

func Sort() {
	//7215441800
	const NUM = 100000000
	var arr []int
	start := time.Now().UnixNano()
	for i := 0; i < NUM; i++ {
		arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		BubbleSort(arr)
	}
	fmt.Println(time.Now().UnixNano() - start)
}

func Sort1() {
	//4242922000
	const NUM = 100000000
	start := time.Now().UnixNano()
	for i := 0; i < NUM; i++ {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		BubbleSort(arr)
	}
	fmt.Println(time.Now().UnixNano() - start)
}

//排序
func BubbleSort(arr []int) {
	for j := 0; j < len(arr)-1; j++ {
		for k := 0; k < len(arr)-1-j; k++ {
			if arr[k] < arr[k+1] {
				temp := arr[k]
				arr[k] = arr[k+1]
				arr[k+1] = temp
			}
		}
	}
}
