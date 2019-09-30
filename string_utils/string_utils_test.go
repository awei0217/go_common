package Persontils

import (
	"fmt"
	"testing"
)

// 字符串拼接的集中方法

func BenchmarkAddStringWithOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddStringWithOperator("时间复杂度和空间复杂度", "s")
	}
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
	t.Log(FindMaxLenNoRepeatSubStr("abcabcbb"))
}
func TestFindMaxLenNoRepeatSubStr2(t *testing.T) {
	t.Log(FindMaxLenNoRepeatSubStr2("abcabcbb"))
}
func TestFindMaxLenNoRepeatSubStr3(t *testing.T) {
	t.Log(FindMaxLenNoRepeatSubStr3("abcad"))
}
func TestFindMaxLenCommonSubStr(t *testing.T) {
	t.Log(FindMaxLenCommonSubStr("gfdef", "abcdef"))
}
func TestFindMaxLenCommonSubStr2(t *testing.T) {
	t.Log(FindMaxLenCommonSubStr2("abc", "abcd"))
}

func TestFindMaxLenCommonSubSeq(t *testing.T) {
	t.Log(FindMaxLenCommonSubSeq("abcd", "anbvcfgd"))
}

func TestFindAndReplacePattern(t *testing.T) {
	t.Log(FindAndReplacePattern([]string{"abb", "dee", "lkj"}, "abb"))
}

func TestRemoveRepeatStr(t *testing.T) {
	t.Log(RemoveRepeatStr("bbaacd"))
}

func TestRepeatedSubstringPattern(t *testing.T) {

	t.Log(RepeatedSubstringPattern("abab"))
}
func TestRepeatedSubstringPattern2(t *testing.T) {

	t.Log(RepeatedSubstringPattern2("ababa"))
}

func TestGetNext(t *testing.T) {
	t.Log(GetNext("abcabc"))

}

func TestStrMatch(t *testing.T) {
	t.Log(StrMatch("adabeabcabc", "abcabc"))
}
func TestIsValid(t *testing.T) {
	t.Log(IsValid("[]}"))

}

func TestLongestPalindrome(t *testing.T) {
	t.Log(LongestPalindrome("cbcdcbedcbc"))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(LengthOfLongestSubstring("abcdan"))
}

func TestDeduplicate2(t *testing.T) {
	t.Log(Deduplicate2("aabcdeer"))
}

func TestLongestCommonPrefix(t *testing.T) {
	t.Log(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func TestRecursionPermutation(t *testing.T) {
	t.Log(RecursionPermutation("acety"))
}

func TestCheckInclusion(t *testing.T) {
	t.Log(CheckInclusion("ab", "mnab"))
}

func TestMultiply(t *testing.T) {
	fmt.Println("1584586728493525893642541066405904799198536104293994395887321573366110169031498781815747478010725266431858021835056717082517535189123595786349457873248" == "1584586728493525893642541066405904799198536104293994395887321573366110169031498781815747478010725266431858021835056717082517535189123595786349457873248")
	t.Log(Multiply("581852037460725882246068583352420736139988952640866685633288423526139", "2723349969536684936041476639043426870967112972397011150925040382981287990380531232"))
}

func Test_LetterCasePermutation(t *testing.T) {
	t.Log(LetterCasePermutation("a1b2"))
}
