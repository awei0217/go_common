package string_utils

import (
	"testing"
)

// 字符串拼接的集中方法

func BenchmarkAddStringWithOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddStringWithOperator("s", "s")
	}
}
func TestAddStringWithOperator(t *testing.T) {

	t.Log(AddStringWithOperator("s", "s"))
}
func BenchmarkAddStringWidthJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddStringWidthJoin([]string{"1", "2"})
	}
}
func BenchmarkAddStringWidthBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddStringWidthBuffer([]string{"1", "2"})
	}
}

func TestReversString(t *testing.T) {
	t.Log(ReversString("spw"))
}

func TestFindMaxLenNoRepeatSubStr(t *testing.T) {
	t.Log(FindMaxLenNoRepeatSubStr("qqwwerrtt"))
}
func TestFindMaxLenCommonSubStr(t *testing.T) {
	t.Log(FindMaxLenCommonSubStr("ab", "abcdef"))
}

func TestFindAndReplacePattern(t *testing.T) {
	t.Log(FindAndReplacePattern([]string{"abb", "dee", "lkj"}, "abb"))
}
