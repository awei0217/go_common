package basic

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"
)

func FileStudy() {
	file, _ := os.Open("E:\\test.txt")

	data := make([]byte, 1024)
	file.Read(data)
	// 验证data字节是否uft-8编码
	fmt.Println(utf8.Valid(data))
	// 从指定位置开始读取多少个字节，以下是从1025 位置开始读取6个字节
	o2, _ := file.Seek(5, 1026)
	b2 := make([]byte, 30)
	file.Read(b2)
	// 字节数组转换string字符串
	fmt.Println(data)
	fmt.Println(b2)
	fmt.Println(o2)

	// 将文件一次性去全部读取到内存
	bs, _ := ioutil.ReadFile("E:\\test.txt")
	fmt.Println(len(bs))

	//获取文件的详细信息
	fileInfo, err := os.Stat("E:\\test.txt")
	//判断文件是否存在
	fmt.Println(os.IsNotExist(err)) //true 不存在 false 存在
	fmt.Println(fileInfo.IsDir())   // 是否是目录 false
	fmt.Println(fileInfo.Mode())    // 文件的权限 -rw-rw-rw-
	fmt.Println(fileInfo.ModTime()) // 最后修改时间
	fileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	fmt.Println(fileSys.CreationTime)
	//获取文件最后写入的时间
	fmt.Println(fileSys.LastWriteTime.Nanoseconds() / 1000 / 1000)
	fmt.Println(time.Unix(fileSys.LastWriteTime.Nanoseconds()/1000/1000/1000, 0).Format("2006-01-02 15:04:05"))
	//获取这个文件的字节大小
	fmt.Println(fileSys.FileSizeLow)
	//按行读取文件
	reader := bufio.NewReader(file)
	for {
		//按行读取文件
		line, err := reader.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil {
			break
		}
		//验证字符串是否是utf-8编码的
		fmt.Println(utf8.ValidString(line))
		fmt.Println(strings.Trim(line, " "))
	}
	createFile, _ := os.Create("E://testgo.txt")
	createFileInfo, _ := createFile.Stat()
	fmt.Println(createFileInfo.IsDir())
	write := bufio.NewWriter(createFile)
	write.WriteString("go 创建的文件\n")
	write.Flush()

	// 创建目录 这样只会创建 spw 目录，同是也创建了test.txt 文件
	os.MkdirAll(path.Dir("E://qwe/test.txt"), os.ModePerm)
	// 这样回把test.txt 创建成一个文件夹
	os.MkdirAll("E://spw/test.txt", os.ModePerm)
	// 创建文件
	spw, _ := os.Create("E://spw/test.txt")
	fmt.Println(spw)

	defer func() {
		file.Close()
		createFile.Close()
		//spw.Close()
	}()

	// 返回是否是绝对路径
	fmt.Println(filepath.IsAbs("E:\\"))
	// 获取绝对路径
	fmt.Println(filepath.Abs("./src"))
	// 返回最后一个路径前面的路径
	fmt.Println(filepath.Dir("E:\\CodeWorkspce\\goProject\\goZiying"))
	//VolumeName函数返回最前面的卷名。如Windows系统里提供参数"C:\foo\bar"会返回"C:"；
	// Unix/linux系统的"\\host\share\foo"会返回"\\host\share"；其他平台会返回""。
	fmt.Println(filepath.VolumeName("E:\\"))
	// 返回文件名的扩展
	fmt.Println(filepath.Ext("E:\\ccjf_price.sql"))

	/**
	遍历某个目录下的所有文件或者文件夹（包含自身）然后调用WalkFunc 函数
	*/
	filepath.Walk("E:\\图片", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, info, err)
		return nil
	})
	// 将路径中的分隔符替换为 斜杠 比如 \\ 替换为 /
	fmt.Println(filepath.ToSlash("E:\\goProject"))
	//FromSlash函数将path中的斜杠（'/'）替换为路径分隔符并返回替换结果，多个斜杠会替换为多个路径分隔符。
	fmt.Println(filepath.FromSlash("E://"))
}
