package basic

import (
	"fmt"
	"reflect"
)

/**
golang  反射学习
*/

type User struct {
	Id   int
	Name string
	Age  int
}

type UserServiceImpl User

type UserService interface {
	/**
	查询
	*/
	FindUserById() User
	/**
	添加
	*/
	InsertUser(user User)
}

func (userService UserServiceImpl) FindUserById() User {
	return User{2, "spw", 25}
}

func (userService UserServiceImpl) InsertUser(u User) {

}

func ReflectionStudy() {
	var num float64 = 1.234
	// 获取num变量的类型
	fmt.Println("type", reflect.TypeOf(num))
	//获取num变量的遏制
	fmt.Println("value", reflect.ValueOf(num))
	//获取num变量的值转化为float64 类型 （转换的时候区分是指针还是类型）
	fmt.Println("value", reflect.ValueOf(num).Interface().(float64))
	// fmt.Println("value",reflect.ValueOf(num).Interface().(*float64)) panic 直接报错，golang 对类型要求非常严格

	/**
	当是未知类型的时候
	*/
	user := User{1, "spw", 25}
	doFiled(user)

	userService := UserServiceImpl(User{})
	doMethod(userService)
}

/**
通过反射获取结构体字段类型和值
*/
func doFiled(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

/**
通过反射获取接口的方法
*/
func doMethod(input interface{}) {
	getValue := reflect.ValueOf(input)
	methodValue := getValue.MethodByName("FindUserById")
	fmt.Println(methodValue.Call(nil)[0])

	getType := reflect.TypeOf(input)
	methodType, flag := getType.MethodByName("FindUserById")
	fmt.Println(methodType, flag)
	fmt.Println(methodType.Name)

}
