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
		AddStringWidthJoin([]string{"时间复杂度和空间复杂度", "2"})
	}
}
func BenchmarkAddStringWidthBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddStringWidthBuffer([]string{"时间复杂度和空间复杂度", "2"})
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
func TestFindMaxLenCommonSubStr2(t *testing.T) {
	t.Log(FindMaxLenCommonSubStr2("gbcdefg", "abcdg"))
}

func TestFindMaxLenCommonSubSeq(t *testing.T) {
	t.Log(FindMaxLenCommonSubSeq("abcdef", "aabcdf"))
}

func TestFindAndReplacePattern(t *testing.T) {
	t.Log(FindAndReplacePattern([]string{"abb", "dee", "lkj"}, "abb"))
}

func TestRemoveRepeatStr(t *testing.T) {
	t.Log(RemoveRepeatStr("aabcd"))
}

func TestRepeatedSubstringPattern(t *testing.T) {

	t.Log(RepeatedSubstringPattern("abab"))
}
func TestRepeatedSubstringPattern2(t *testing.T) {

	t.Log(RepeatedSubstringPattern2("ababa"))
}

func TestGetNext(t *testing.T) {
	GetNext("ababac")
}

func TestStrMatch(t *testing.T) {
	t.Log(StrMatch("abababdacabcda", "ababda"))
}
func TestIsValid(t *testing.T) {
	t.Log(IsValid("[]}"))
}