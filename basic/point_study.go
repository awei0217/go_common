package basic

import (
	"fmt"
)

func PointStudy() {
	var pointInt1 = 1

	fmt.Println(pointInt1)
	//Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
	fmt.Println(&pointInt1)

	// 使用指针
	var pointInt2 *int // 声明一个指针
	var pointInt3 = 1  //声明一个实际变量
	pointInt2 = &pointInt3

	// & 打印出来的是内存地址  * 打印出来的是值
	fmt.Println(pointInt2)
	fmt.Println(*pointInt2)

	// 指针数组
	var arrayPointInt1 [10]*int
	for i := 0; i < len(arrayPointInt1); i++ {
		arrayPointInt1[i] = &i //给数组赋值
	}
	fmt.Println(arrayPointInt1)

	for _, value := range arrayPointInt1 {
		fmt.Println(*value) //打印数组值
	}
	var x, y *int
	var c, d = 1, 1
	x, y = &c, &d
	pointMethod(x, y)

}

// 指针函数

func pointMethod(x, y *int) {
	fmt.Println(x == nil)
	fmt.Println(y == nil)
	fmt.Println(*x + *y)
}

func SlicePointStudy() {
	slice := make([]int, 0, 0)
	slice = slicePoint(slice)
	fmt.Println(slice)
}
func slicePoint(slice []int) []int {
	for i := 0; i < 5; i++ {
		slicePoint(slice)
	}
	return slice

}
