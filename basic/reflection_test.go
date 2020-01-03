package basic

import (
	"fmt"
	"testing"
)

func TestReflectionStudy(t *testing.T) {
	ReflectionStudy()
}

func TestFuncReflection(t *testing.T) {
	FuncReflection(func(s string) {
		fmt.Println(s)
	}, "ss")
}

func TestReflect01(t *testing.T) {
	var num int64 = 10
	Reflect01(num)
}
