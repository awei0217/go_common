package algorithm

import (
	"bytes"
	"fmt"
	"strings"
)

// 字符串拼接的集中方法


func AddStringWithOperator(){
	hello := "hello"
	world := "world"
	_ = hello + world
}

func AddStringWidthSprintf()  {
	hello := "hello"
	world := "world"
	_ = fmt.Sprintf("%s,%s",hello,world)
}

func AddStringWidthJoin()  {
	hello := "hello"
	world := "world"
	_ = strings.Join([]string{hello,world},"")
}

func AddStringWidthBuffer()  {
	hello := "hello"
	world := "world"
	var buffer bytes.Buffer
	buffer.WriteString(hello)
	buffer.WriteString(world)
	_ = buffer.String()
}