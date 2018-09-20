package string_utils

import (
	"bytes"
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
	start1 :=-1
	start2 :=-1
	comparisons := 0
	longest := 0
	for i:=0;i<len(str1);i++{
		for j:=0;j<len(str2);j++{
			length:=0
			m:=i
			n:=j
			for m<len(str1) && n<len(str2) {
				comparisons++
				if str1[m] != str2[n]{
					break
				}
				length++
				m++
				n++
			}
			if longest < length{
				longest = length
				start1 =i
				start2 =j
			}
		}
	}
	if len(str1) > len(str2){
		return str1[start1:start1+longest]
	}else{
		return str1[start2:start2+longest]
	}
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
