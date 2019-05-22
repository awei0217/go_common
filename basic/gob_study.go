package basic

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

/**
各个类型的编解码规则
整型：分为sign int和usign int， 其中从上面例子也看到，int和uint是不能互相编解码的。float和int也是不能互相编解码的。
Struct，array，slice是可以被编码的。但是function和channel是不能被编码的。
bool类型是被当作uint来编码的，0是false，1是true。
浮点类型的值都是被当作float64类型的值来编码的
String和[]byte传递是uint(byte个数) + byte[]的形式编码的
Slice和array是按照uint(array个数) + 每个array编码 这样的形式进行编码的
Maps是按照 uint(Map个数) + 键值对 这样的形式进行编码的
*/
type P struct {
	X int
	Y string
}
type Q struct {
	Y string
	X int
	Z string
}

/**
字段先后顺序不影响编解码
*/
func EncoderAndDecoder(p P, q Q) {
	buf := bytes.Buffer{}
	gob.NewEncoder(&buf).Encode(p)
	gob.NewDecoder(&buf).Decode(&q)
	fmt.Println(q)

}
