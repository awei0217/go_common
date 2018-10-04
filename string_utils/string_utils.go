package string_utils

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

/**
通过+号拼接字符串
*/
func AddStringWithOperator(str1, str2 string) string {

	return str1 + str2
}

/**
通过strings 包的join连接字符串
*/
func AddStringWidthJoin(strArray []string) string {
	return strings.Join(strArray, "")
}

/**
通过buffer 拼接字符串
*/
func AddStringWidthBuffer(strArray []string) string {
	var buffer bytes.Buffer
	for _, value := range strArray {
		buffer.WriteString(value)
	}
	return buffer.String()
}

/**
反转字符串
*/
func ReversString(str string) string {
	count := len(str)
	bytes := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		bytes[i] = str[count-1-i]
	}
	return string(bytes)
}

/**
查找给定字符串中的最长不重复子串
返回最长不重复子串+子串的长度
*/
func FindMaxLenNoRepeatSubStr(s string) (string, int) {
	head, tail := 0, 0
	maxLenNoRepeatSubStr := ""
	for i := 0; i < len(s)-1; i++ {
		for j := i; j < len(s); j++ {
			if strings.Contains(maxLenNoRepeatSubStr, s[j:j+1]) {
				if head == 0 && tail == 0 {
					head, tail = i, j
				}
				if len(s[i:j]) > len(s[head:tail]) {
					head, tail = i, j
				}
				maxLenNoRepeatSubStr = ""
				break
			}
			maxLenNoRepeatSubStr = s[i : j+1]
		}
		if maxLenNoRepeatSubStr == s[i:] {
			head, tail = i, len(s)
			break
		}
	}
	return s[head:tail], tail - head
}

/**
从两个给定字符串中找出最长公共子串
返回最长公共子串
*/
func FindMaxLenCommonSubStr(str1, str2 string) string {
	start1 := -1
	start2 := -1
	comparisons := 0
	longest := 0
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			length := 0
			m := i
			n := j
			for m < len(str1) && n < len(str2) {
				comparisons++
				if str1[m] != str2[n] {
					break
				}
				length++
				m++
				n++
			}
			if longest < length {
				longest = length
				start1 = i
				start2 = j
			}
		}
	}
	if len(str1) > len(str2) {
		return str1[start1 : start1+longest]
	} else {
		return str1[start2 : start2+longest]
	}
}

// 采用动态规划求取
func FindMaxLenCommonSubStr2(str1, str2 string) string {
	l1 := len(str1)
	l2 := len(str2)
	max := 0
	end := 0
	var twoArray [][]int
	for i := 0; i < l1+1; i++ {
		tmp := make([]int, l2+1)
		twoArray = append(twoArray, tmp)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if str1[i-1] == str2[j-1] {
				twoArray[i][j] = twoArray[i-1][j-1] + 1
				if twoArray[i][j] > max {
					max = twoArray[i][j]
					end = j
				}
			} else {
				twoArray[i][j] = 0
			}
		}
	}
	fmt.Println(twoArray)
	bytes := make([]byte, 0)
	for m := end - max; m < end; m++ {
		bytes = append(bytes, str2[m])
	}
	return string(bytes)
}

/**
求最长公共子序列
*/
func FindMaxLenCommonSubSeq(str1, str2 string) string {
	l1 := len(str1)
	l2 := len(str2)
	var twoArray [][]int
	for i := 0; i < l2+1; i++ {
		tmp := make([]int, l1+1)
		twoArray = append(twoArray, tmp)
	}
	bs := make([]byte, 0)
	for m := 1; m <= l2; m++ {
		for n := 1; n <= l1; n++ {
			if str1[n-1] == str2[m-1] {
				twoArray[m][n] = twoArray[m-1][n-1] + 1
				bs = append(bs, str1[n-1])
			} else if twoArray[m-1][n] >= twoArray[m][n-1] {
				twoArray[m][n] = twoArray[m-1][n]
			} else {
				twoArray[m][n] = twoArray[m][n-1]
			}
		}
	}
	return RemoveRepeatStr(string(bs))
}

/**
移除字符串中的重复字符
*/
func RemoveRepeatStr(str string) string {
	if len(str) == 0 {
		return ""
	}
	bs := [256]byte{}
	for _, v := range str {
		bs[v] = 1
	}
	rs := make([]byte, 0)
	for index, v := range bs {
		if v == 1 {
			rs = append(rs, byte(index))
		}
	}
	return string(rs)
}

/**
你有一个单词列表 words 和一个模式  pattern，你想知道 words 中的哪些单词与模式匹配。
如果存在字母的排列 p ，使得将模式中的每个字母 x 替换为 p(x) 之后，我们就得到了所需的单词，那么单词与模式是匹配的。
（回想一下，字母的排列是从字母到字母的双射：每个字母映射到另一个字母，没有两个字母映射到同一个字母。）
返回 words 中与给定模式匹配的单词列表。
你可以按任何顺序返回答案。
*/
func FindAndReplacePattern(words []string, pattern string) []string {
	patternWords := make([]string, 0)
	for _, word := range words {
		flag := true
		ruleMap1 := make(map[byte]byte, len(pattern))
		ruleMap2 := make(map[byte]byte, len(pattern))
		for j := 0; j < len(pattern); j++ {
			p := pattern[j]
			w := word[j]
			if _, ok := ruleMap1[p]; ok {
				if ruleMap1[p] != w {
					flag = false
					break
				}
			} else if _, ok := ruleMap2[w]; ok {
				flag = false
			} else {
				ruleMap1[p] = w
				ruleMap2[w] = p
			}
		}
		if flag {
			patternWords = append(patternWords, word)
		}
	}
	return patternWords
}

/**
判断字符串是否为空
true 为空 false 不为空
*/
func IsBlank(str string) bool {
	return !(len(str) > 0)
}

/**
给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000
"abab" true   "aba" false
*/
func RepeatedSubstringPattern(s string) bool {
	length := len(s)
	if length == 0 || length == 1 {
		return false
	}
	n := 2
	for n <= length {
		mid := length / n
		step := mid
		index := 0
		flag := true
		for mid < length {
			if (mid+step) > length || s[index:mid] != s[mid:mid+step] {
				flag = false
				break
			}
			index = index + step
			mid = mid + step
		}
		if flag {
			fmt.Println(step)
			return true
		}
		n++
	}
	return false
}
func RepeatedSubstringPattern2(s string) bool {
	if len(s) == 0 {
		return false
	}
	size := len(s)
	ss := (s + s)[1 : size*2-1]
	return strings.Contains(ss, s)
}

func GetNext(p string) []int { //ababda
	next := make([]int, 0)
	next = append(next, 0)
	k := 0
	for i := 1; i < len(p); i++ {
		for k > 0 && p[i] != p[k] {
			k = next[k-1]
		}
		if p[i] == p[k] {
			k = k + 1
		}
		next = append(next, k)
	}
	return next
}

/**
KMP 算法，字符串模式匹配算法 主要是 GetNext
*/
func StrMatch(source, target string) int {
	slen := len(source)
	tlen := len(target)
	next := GetNext(target)
	q := 0
	for i := 0; i < slen; i++ {
		for q > 0 && target[q] != source[i] {
			q = next[q-1]
		}
		if target[q] == source[i] {
			q++
		}
		if q == tlen {
			return i - tlen + 1
		}
	}
	return -1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var buf = bytes.Buffer{}

func Tree2Str(t *TreeNode) string {
	left := 0
	if t != nil {
		buf.WriteString(strconv.Itoa(t.Val))
		if t.Left != nil {
			left = 1
			buf.WriteString("(")
			Tree2Str(t.Left)
			buf.WriteString(")")
		}
		if t.Right != nil {
			if left == 0 {
				buf.WriteString("()")
			}
			buf.WriteString("(")
			Tree2Str(t.Right)
			buf.WriteString(")")
		}
	}
	return buf.String()
}
