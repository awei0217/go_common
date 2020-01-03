package basic

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

type MyInt int
type I MyInt
type Ints []int
type MyInts []MyInt
type M map[string]string
type CustomM M

type aliasInt64 = int64

func GetDiCengType() {
	var q MyInt = 1
	var w I = 2
	var e Ints = []int{1, 2, 3}
	var t M = make(map[string]string)

	fmt.Println(reflect.TypeOf(q).Kind())
	fmt.Println(reflect.TypeOf(w).Kind())
	fmt.Println(reflect.TypeOf(e).Kind())
	fmt.Println(reflect.TypeOf(t).Kind())

	var r rune = 'c'
	var b byte = 1
	fmt.Println(reflect.TypeOf(r).Kind())
	fmt.Println(reflect.TypeOf(b).Kind())

}

func CustomType() {
	var i MyInt = 2
	var j int = 1

	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(j))
	fmt.Println(reflect.TypeOf(i).Kind(), reflect.TypeOf(j).Kind())

}

type MyM struct {
	i int64
}
type MyN struct {
	i int64
}

func TestStruct() {
	n := MyN{i: 10}
	var m MyM
	m = MyM(n)
	fmt.Println(n, m)
}

//
//fmt.Println(reflect.TypeOf(m), reflect.TypeOf(n))
//fmt.Println(reflect.TypeOf(m).Kind(), reflect.TypeOf(n).Kind())
func ByteToString() {
	var s byte = 1
	fmt.Println(string(s))
}

func RuneToString() {
	var r rune = 'A'

	fmt.Println(string(r))

}

func StringToRune() {
	var s string = "AB"

	fmt.Println([]rune(s))
}

func StringToByte() {
	var s string = "AB"

	fmt.Println([]byte(s))
}

func StringTypeToRuneToByte() {
	var s string = "AB"
	fmt.Println(reflect.TypeOf(s[0]))
	for _, v := range s {
		fmt.Println(reflect.TypeOf(v)) //int32类型
		fmt.Println(string(v))
	}
}

func A1() {

	var v interface{} = 1
	var s uint8 = 1

	temp1 := int(s)
	temp2 := v.(int)

	fmt.Println(temp2)
	fmt.Println(temp1)

}

func ConvertType() {
	var i int8 = 123
	var j int16 = int16(i)

	var m int64 = 128
	var n int8 = int8(m)

	fmt.Println(j, n)
}

func StrConvertType() {
	var s1, s2 string = "AbcD", "1234"

	//转字节
	bs1 := []byte(s1)
	bs2 := []byte(s2)
	//字节数组转字符串
	s11 := string(bs1)
	s22 := string(bs2)
	//单个字节转字符串
	ss := string(bs1[0])
	fmt.Println(s11, s22, ss)

	//s2转数字 ,err 表示是否能转换成功，比如s1就会转换失败
	i, err := strconv.Atoi(s2)
	//数字转字符串
	s := strconv.Itoa(i)

	//字符串转字符数组
	runes := []rune(s1)

	//字符数组转字符串
	ss1 := string(runes)
	//单个字符转字符串
	ss2 := strconv.QuoteRune(runes[0])

	//字符转字节
	bss := make([]byte, 0)
	bss = strconv.AppendQuoteRune(bss, runes[0])

	fmt.Println(err, s, ss1, ss2, runes[0], bss, string(bss))
	//除开rune和byte底层的类型的区别，在使用上，rune能处理一切的字符，而byte仅仅局限在ascii

	//整形转字节
	x := int32(68)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	//字节转整形
	var y int32
	binary.Read(bytesBuffer, binary.BigEndian, &y)
}

func InterfaceConvertType() {
	var i interface{} = 1
	t, f := i.(int)
	if f {
		fmt.Println(t)
	} else {
		fmt.Println(reflect.TypeOf(i).Kind(), reflect.TypeOf(i))
	}
}

//uintptr 是一个整数类型，用来表示任意地址,其大小足以容纳任何指针的位模式，可以用来进行数值计算
//unsafe.Pointer是一个指针类型。unsafe.Pointer值不能被取消引用。
// 如果unsafe.Pointer变量仍然有效，则由unsafe.Pointer变量表示的地址处的数据不会被GC回收
func UintptrAndUnsafe() {

	var p uintptr = 9223372036854775807*2 + 1

	fmt.Println(p)

	var s string = "1"

	var i uintptr = 1

	a1 := unsafe.Pointer(i)
	a2 := unsafe.Pointer(&i)
	a3 := unsafe.Pointer(&s)
	a4 := uintptr(a1)

	fmt.Println(s, i, a1, a2, a3, uintptr(a1), a4)

	m := unsafe.Sizeof(a1)

	fmt.Println(m)
	//================================================
	a := [4]int{0, 1, 2, 3}
	p1 := unsafe.Pointer(&a[1])
	p3 := unsafe.Pointer(uintptr(p1) + 2*unsafe.Sizeof(a[0]))
	*(*int)(p3) = 6
	fmt.Println("a =", a) // a = [0 1 2 6]

	// ...

	type Person struct {
		name   string
		age    int
		gender bool
	}

	who := Person{"John", 30, true}
	pp := unsafe.Pointer(&who)
	pname := (*string)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.name)))
	page := (*int)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.age)))
	pgender := (*bool)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.gender)))
	*pname = "Alice"
	*page = 28
	*pgender = false
	fmt.Println(who) // {Alice 28 false}

	type W struct {
		b int32
		c int64
	}
	var w *W = new(W)
	fmt.Println(w.b)
	fmt.Println(unsafe.Alignof(w.b))
	fmt.Printf("size=%d\n", unsafe.Sizeof(*w))

	pb := (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b)))
	*pb = 32
	fmt.Println(w.b)

	//现在我们通过指针运算给b变量赋值为10
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	*((*int)(b)) = 10
	//此时结果就变成了10，0
	fmt.Println(w.b, w.c)

	fmt.Println(unsafe.Sizeof(who))

	slice1 := make([]int8, 10)
	fmt.Println(unsafe.Sizeof(&slice1))

}

func UintptrAndPointerStudty() {
	//创建一个变量
	var i int8 = 10

	//建一个变量转化成Pointer 和 uintptr
	p := unsafe.Pointer(unsafe.Sizeof(i)) //入参必须是指针类型的
	fmt.Println(p)
	u := unsafe.Sizeof(i)
	fmt.Println(u)
	//uintptr和Pointer转换
	u = unsafe.Sizeof(p) //传入指针，获取的是指针的大小
	fmt.Println(u)
	u = unsafe.Sizeof(i) //获取的是变量的大小
	fmt.Println(u, "ss")
	fmt.Println(uintptr(i))

	p = unsafe.Pointer(u)
	fmt.Println(p) // 1

	//创建一个结构体
	type Person struct {
		b int64
		a bool
		c int8
		d string
	}
	//接下来演示一下内存对齐
	person := Person{a: true, b: 1, c: 1, d: "spw"}
	fmt.Println(unsafe.Alignof(person))

	type Person2 struct {
		a bool
		c int8
	}

	p2 := Person2{a: true, c: 1}
	fmt.Println(unsafe.Alignof(p2))
	type Mutex struct {
		state int32
		sema  uint32
	}
	type poolLocalInternal struct {
		private interface{}   // Can be used only by the respective P.
		shared  []interface{} // Can be used by any P.
		Mutex                 // Protects shared.
	}

	fmt.Println(unsafe.Sizeof(poolLocalInternal{}))
}
