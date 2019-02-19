package basic

import (
	"bufio"
	"fmt"
	"os"
	"text/scanner"
	"unicode/utf8"
)


func ScannerStudy(){

	var s scanner.Scanner
	file,_ := os.Open("E:\\新建文本文档.txt")
	s.Init(bufio.NewReader(file))
	tok := s.Scan()
	fmt.Println(tok)
	//验证是不是utf8的字符
	fmt.Println(utf8.ValidRune(tok))

}
