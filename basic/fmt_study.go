package basic

import (
	"fmt"
	"strings"
)

// fmt包学习

func FmtStudy() {
	//支持可变参数输入，多个用，隔开；转化为字符串标准输出，返回写入的字节数
	// fmt.Print(a...interface()) (n int, err error)

	//跟fmt.Print()一样，只是后面加了个换行符
	//fmt.Println(a...interface())(n int, err error)

	//Printf 将字符串填充到前面的占位符中，返回写入的字节数
	//fmt.Printf(format string, a ...interface{})(n int, err error)

	// 功能同上面三个函数，只不过将转换结果写入到 流中。
	//func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	//func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	//func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)

	// 功能同上面三个函数，只不过将转换结果以字符串形式返回。
	//func Sprint(a ...interface{}) string
	//func Sprintln(a ...interface{}) string
	//func Sprintf(format string, a ...interface{}) string

	// 功能同 Sprintf，只不过结果字符串被包装成了 error 类型。
	//func Errorf(format string, a ...interface{}) error

	// Scan 从标准输入中读取数据，并将数据用空白分割并解析后存入 a 提供
	// 的变量中（换行符会被当作空白处理），变量必须以指针传入。
	// 当读到 EOF 或所有变量都填写完毕则停止扫描。
	// 返回成功解析的参数数量。
	//func Scan(a ...interface{}) (n int, err error)
	i, str := 12, "123"
	count, _ := fmt.Scan(&i, &str)
	fmt.Printf("%d", count)
	fmt.Println()

	// Scanln 和 Scan 类似，只不过遇到换行符就停止扫描。
	//func Scanln(a ...interface{}) (n int, err error)

	// Scanf 从标准输入中读取数据，并根据格式字符串 format 对数据进行解析，
	// 将解析结果存入参数 a 所提供的变量中，变量必须以指针传入。
	// 输入端的换行符必须和 format 中的换行符相对应（如果格式字符串中有换行
	// 符，则输入端必须输入相应的换行符）。
	// 占位符 %c 总是匹配下一个字符，包括空白，比如空格符、制表符、换行符。
	// 返回成功解析的参数数量。
	//func Scanf(format string, a ...interface{}) (n int, err error)

	// 功能同上面三个函数，只不过从 r 中读取数据。
	//func Fscan(r io.Reader, a ...interface{}) (n int, err error)
	//func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
	//func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
	count, _ = fmt.Fscan(strings.NewReader("从流中读取数据"), str)
	fmt.Println(count)

	// 功能同上面三个函数，只不过从 str 中读取数据。
	//func Sscan(str string, a ...interface{}) (n int, err error)
	//func Sscanln(str string, a ...interface{}) (n int, err error)
	//func Sscanf(str string, format string, a ...interface{}) (n int, err error)

	//通用==============================================================

	// %v 以默认格式打印
	fmt.Printf("%v", "这是一个字符串")
	fmt.Println()
	// %+v,类似与%v,但输出结构体时会添加字段名
	fmt.Printf("%+v", &Student{Name: "spw", Age: 15})
	fmt.Println()
	fmt.Printf("%v", &Student{Name: "spw", Age: 15})
	fmt.Println()
	// %#v值的go语法表示
	fmt.Printf("%#v", &Student{Name: "spw", Age: 15})
	fmt.Println()
	//%T值的类型的Go语法表示
	fmt.Printf("%T", &Student{Name: "spw", Age: 15})
	fmt.Println()
	//%% 表示百分号
	fmt.Printf("%%")
	fmt.Println()

	//================================================================
	//单词true或者false
	fmt.Printf("%t", true)
	fmt.Printf("%t", false)
	fmt.Println()

	//===============================================================
	/*%b	表示为二进制
	%c	该值对应的unicode码值
	%d	表示为十进制
	%o	表示为八进制
	%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	%x	表示为十六进制，使用a-f
	%X	表示为十六进制，使用A-F
	%U	表示为Unicode格式：U+1234，等价于"U+%04X"*/
	fmt.Printf("%b", 11)
	fmt.Println()
	fmt.Printf("%c", 12)
	fmt.Println()
	fmt.Printf("%d", 11)
	fmt.Println()
	fmt.Printf("%o", 11)
	fmt.Println()
	fmt.Printf("%q", 11)
	fmt.Println()
	fmt.Printf("%x", 11)
	fmt.Println()
	fmt.Printf("%X", 11)
	fmt.Println()
	fmt.Printf("%U", 11)
	fmt.Println()

	//======================================================================
	//浮点数与复数的两个组分：
	/**
	%b	无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
	%e	科学计数法，如-1234.456e+78
	%E	科学计数法，如-1234.456E+78
	%f	有小数部分但无指数部分，如123.456
	%F	等价于%f
	%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
	*/
	fmt.Printf("%b", 12.12)
	fmt.Println()
	fmt.Printf("%e", 12.12)
	fmt.Println()

	//======================================================================================
	//字符串和[]byte：
	/**
	%s	直接输出字符串或者[]byte
	%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	%x	每个字节用两字符十六进制数表示（使用a-f）
	%X	每个字节用两字符十六进制数表示（使用A-F）
	*/

}
