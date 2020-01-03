package basic

import (
	"unicode"
)

//是否是字母
func IsLetter(r rune) bool {

	return unicode.IsLetter(r)
}

//是否是数字
func IsDigit(r rune) bool {

	return unicode.IsDigit(r)
}

//是否特殊字符
func IsTitle(r rune) bool {

	return unicode.IsTitle(r)
}
