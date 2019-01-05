package basic

import (
	"fmt"
	"strconv"
)

func StrconvStudy() {
	//ParseBoll 将字符串转换为布尔值
	// 它接受真值：1, t, T, TRUE, true, True
	// 它接受假值：0, f, F, FALSE, false, False.
	// 其它任何值都返回一个错误
	//func ParseBool(str string) (value bool, err error)
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("0"))
	// FormatBool 将布尔值转换为字符串 "true" 或 "false"
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatBool(false))
	// ParseFloat 将字符串转换为浮点数
	//s：要转换的字符串
	// bitSize：指定浮点类型（32:float32、64:float64）
	// 如果 s 是合法的格式，而且接近一个浮点值，
	// 则返回浮点数的四舍五入值（依据 IEEE754 的四舍五入标准）
	// 如果 s 不是合法的格式，则返回“语法错误”
	// 如果转换结果超出 bitSize 范围，则返回“超出范围”
	fmt.Println(strconv.ParseFloat("2.333", 32))
	fmt.Println(strconv.ParseFloat("2.333", 64))

	// ParseInt 将字符串转换为 int 类型
	// s：要转换的字符串
	// base：进位制（2 进制到 36 进制）
	// bitSize：指定整数类型（0:int、8:int8、16:int16、32:int32、64:int64）
	// 返回转换后的结果和转换时遇到的错误
	// 如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）
	fmt.Println(strconv.ParseInt("20", 10, 0))
	fmt.Println(strconv.ParseInt("0xFFF", 0, 0))
	// 将字符串转化为int
	fmt.Println(strconv.Atoi("1"))
	fmt.Println(strconv.Atoi("-5"))

	// FormatUint 将 int 型整数 i 转换为字符串形式
	// base：进位制（2 进制到 36 进制）
	// 大于 10 进制的数，返回值使用小写字母 'a' 到 'z'
	//func FormatInt(i int64, base int) string
	fmt.Println(strconv.FormatInt(1, 10))
	// 将int 转化为字符串
	fmt.Println(strconv.Itoa(-3))

	// FormatFloat 将浮点数 f 转换为字符串值
	// f：要转换的浮点数
	// fmt：格式标记（b、e、E、f、g、G）
	// prec：精度（数字部分的长度，不包括指数部分）
	// bitSize：指定浮点类型（32:float32、64:float64）
	// 格式标记：
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)
	// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
	// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）
	//func FormatFloat(f float64, fmt byte, prec, bitSize int) string

	fmt.Println(strconv.FormatFloat(2.44, 'f', 10, 64)) //2.4400000000
	fmt.Println(strconv.FormatFloat(2.44, 'g', 10, 64)) //2.44
	fmt.Println(strconv.FormatFloat(2.44, 'G', 10, 64)) //2.44

}
