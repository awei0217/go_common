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
