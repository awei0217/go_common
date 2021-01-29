package basic

import (
	"fmt"
)

/**
执行顺序，先声明的后执行
*/
func DeferStudy() {
	defer fmt.Println("defer 1")
	defer func() {
		fmt.Println("defer 2")
	}()
	defer func(str string) {
		fmt.Println(str)
	}(string("带入参的defer 3"))
}

func RecoverStudy() error {
	/**
	defer 必须声明到panic之前
	*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕捉到panic:", err)
		}
	}()
	panic("模拟发生错误")

}

func RecoverA() (i int, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕捉到panic:", r)
			err = r.(error)
		}
	}()
	panicA()
	return 0, err
}

func panicA() {
	bs := make([]byte, 4, 4)
	bs[10] = 0
}

/**
先来假设出结论，帮助大家理解原因：
多个defer的执行顺序为“后进先出”；
defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。

如何解释两种结果的不同：
上面两段代码的返回结果之所以不同，其实从上面第2条结论很好理解。
a()int 函数的返回值没有被提前声名，其值来自于其他变量的赋值，而defer中修改的也是其他变量，而非返回值本身，因此函数退出时返回值并没有被改变。
b()(i int) 函数的返回值被提前声名，也就意味着defer中是可以调用到真实返回值的，因此defer在return赋值返回值 i 之后，再一次地修改了 i 的值，最终函数退出后的返回值才会是defer修改过的值
*/
func A() int { //返回0
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}

func B() (i int) { //返回2
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}

//测试defer性能和正常关闭性能
type channel chan int

func NoDefer() {
	ch1 := make(channel, 10)
	close(ch1)
}
func Defer() {
	ch2 := make(channel, 10)
	defer close(ch2)
}
