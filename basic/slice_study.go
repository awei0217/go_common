package basic

import "fmt"

//Go 语言切片是对数组的抽象。

//Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，
//Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

// 声明切片 （未指定大小的数组）
var slice1 []int

func SliceStudy() {
	// 使用make 函数创建切片,第一个参数为类型，第二个参数为长度length，第三个参数为容量 capacity
	slice1 = make([]int, 10, 20)

	//切片初始化
	slice2 := []int{1, 2, 3}
	// len函数获取切片的长度  cap 函数获取切片的容量
	fmt.Println(slice2, len(slice2), cap(slice2))

	var arraySliceInt = [...]int{1, 2, 3, 4, 5, 6}

	slice3 := arraySliceInt[:]   // 将数组赋值给切片
	slice4 := arraySliceInt[0:2] //数组下标0到2的值赋值给切片 ,不包含末尾
	slice5 := arraySliceInt[0:]  //数组下标0到结尾赋值给切片
	slice6 := arraySliceInt[:5]  //数组下标0开始到5赋值给切片,  不包含末尾
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)

	slice7 := slice3[:] // 通过切片slice3初始化切片slice7
	fmt.Println(slice7)

	//切面的操作
	fmt.Println(slice7[1:5]) //打印子切面

	slice8 := append(slice7, 10) // append 方法给切面增加元素
	fmt.Println(slice8)
	var slice9 = make([]int, 20, len(slice8)*2) // 创建切片slice9 是slice8切面容量的两倍
	len := copy(slice9, slice8)                 // copy 方法把slice8  复制到 slice9,返会slice9 切面的长度
	fmt.Println(len)

	var ss []string
	print("func print", ss)
	//切片尾部追加元素append elemnt
	for i := 0; i < 10; i++ {
		ss = append(ss, fmt.Sprintf("s%d", i))
	}
	print("after append", ss)
	//删除切片元素remove element at index
	index := 5
	ss = append(ss[:index], ss[index+1:]...)
	print("after delete", ss)
	//在切片中间插入元素insert element at index;
	//注意：保存后部剩余元素，必须新建一个临时切片
	rear := append([]string{}, ss[index:]...)
	ss = append(ss[0:index], "inserted")
	ss = append(ss, rear...)
	print("after insert", ss)
}

func S() (int, int) {
	return 1, 1
}
