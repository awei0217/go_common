package main

import (
	"fmt"
)

/**
可以写两个init 方法，按顺序执行
*/
func init() {
	fmt.Println("A")
}

func init() {
	fmt.Println("B")
}

func main() {

}
