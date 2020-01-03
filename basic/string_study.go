package basic

import (
	"bytes"
	"fmt"
	"strings"
)

//Go语言支持以下2种形式的字符串：
//1. 解释性字符串：带引号的字节序列。该类字符串使用双引号括起来，其中的相关的转义字符将被替换。例如：
//str ：= "laoYu"
//2. 原生字符串： 该类字符串使用反引号（注意：不是单引号）括起来，支持换行。例如：
//str:=`This is a raw string \n`
//上面原生字符串中的 \n 会被原样输出。

func StringStudy() {

	str1 := "hello\ngo"

	fmt.Println(str1) // 会换行
	str2 := `hello \t go \n`

	fmt.Println(str2) //会原样输出

	str3 := "sunpengwei"
	//字符串函数包
	fmt.Println(len(str3)) //打印字符串长度

	fmt.Println(strings.HasPrefix(str3, "s"))    // 判断是否以 s 开头
	fmt.Println(strings.HasSuffix(str3, "i"))    // 判断是否以 i 结尾
	fmt.Println(strings.Contains(str3, "peng"))  // 判断是否包含"peng"字符串
	fmt.Println(strings.ContainsAny(str3, "hy")) // 只要原始字符串中包含子字符串中的任意一个字符，都返回true,否则返回false
	// 比较两个字符串的大小，一模一样 返回0,str3 小于"www" 返回 -1，大于的话返回 1
	fmt.Println(strings.Compare(str3, "www"))
	fmt.Println(strings.Count(str3, "n"))              // 统计字符串出现的次数
	fmt.Println(strings.EqualFold(str3, "sunpengwei")) // 比较两个字符串是否相等
	fmt.Println(strings.Fields(str3))                  //将字符串转换为一个切片
	fmt.Println(strings.Split("s,p,w", ","))           // 根据子字符串切割，返回一个切片
	fmt.Println(strings.Repeat(str3, 5))               // 把字符串连接n次拼成一个新的字符串
	// 把字符串wei 替换成 yyy 0代表替换0次，-1 代表所有的wei都替换，1,2,3 分别代表替换几次
	fmt.Println(strings.Replace(str3, "wei", "yyy", -1))
	// 根据子字符串切割，每个字符串保留分隔符（除了最后一个），返回一个切片
	fmt.Println(strings.SplitAfter("s,pp,w", ","))
	fmt.Println(strings.Index(str3, "n"))      // 返回字符串第一次出现的索引
	fmt.Println(strings.IndexAny(str3, "hnw")) // 返回字符串第一个被原始字符串包含的字符第一次出现的索引
	//使用空格连接字符串
	var s = make([]string, 0, 0)
	s = append(s, "sun", "peng")
	fmt.Println("Join", strings.Join(s, ""))

	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))     // ["Achtung"]
	fmt.Printf("[%q]", strings.TrimLeft(" !!! Achtung !!! ", "! ")) // ["Achtung !!! "]
	fmt.Println(strings.TrimSpace("             a lone gopher   ")) // a lone gopher

	// 拼接字符串比 + 速度快
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	for i := 0; i < 1000000; i++ {
		/*
		   func (b *Buffer) WriteString(s string) (n int, err error)
		   Write将s的内容写入缓冲中，如必要会增加缓冲容量。返回值n为len(p)，err总是nil。如果缓冲变得太大，Write会采用错误值ErrTooLarge引发panic。
		*/
		buffer.WriteString("s")
	}

	fmt.Println("拼接后的结果为-->", buffer.String())

	fmt.Println(string([]rune("2018-08-10")[:7]))
}

var hm = make(map[string]string, 8)
var result = make([]string, 0)

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	hm["2"] = "abc"
	hm["3"] = "def"
	hm["4"] = "ghi"
	hm["5"] = "jkl"
	hm["6"] = "mno"
	hm["7"] = "pqrs"
	hm["8"] = "tuv"
	hm["9"] = "wxyz"

	backtrack("", digits)

	fmt.Println(result)
	return result
}

func backtrack(str, next_digits string) {
	if len(next_digits) == 0 {
		result = append(result, str)
	} else {
		v := next_digits[:1]
		temp := hm[v]
		for _, j := range temp {
			leter := j
			backtrack(str+string(leter), next_digits[1:])
		}
	}
}
