package basic

import "fmt"

/**
go语言map学习
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值
*/

// 可以使用make函数 或者 map关键字 来定义map

//声明一个map
var mapInt1 map[int]string

//使用make函数创建map
var mapInt2 = make(map[int]string)

func MapStudy() {
	mapInt1 = make(map[int]string, 10)
	fmt.Println(len(mapInt1)) //0

	// 为map插入键值对
	mapInt1[0] = "spw"
	mapInt1[1] = "wei"
	mapInt1[2] = "yyy"

	// 使用key输出map的值
	for key := range mapInt1 {
		fmt.Printf("map key %d value is %s\n", key, mapInt1[key])
	}

	// 检查集合中元素是否存在
	value, ok := mapInt1[3]
	if ok {
		fmt.Println("map key 3 is " + value)
	} else {
		fmt.Println("map key 3 is null")
	}

	// delete 函数 ，用于删除集合元素 参数为map 和 key

	delete(mapInt1, 0)
	// 删除之后打印map
	for key, value := range mapInt1 {
		fmt.Printf("map key %d value is %s\n", key, value)
	}

}
