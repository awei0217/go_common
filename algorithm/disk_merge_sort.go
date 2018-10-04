package algorithm

import (
	"bufio"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)
// 磁盘文件归并排序
const (
	//每次读取多少行写排序后写入文件
	NUMBER_TO_SORT = 1000
)
/**
	读取大文件排序写入临时文件
 */
func RedFileSortWriteTempFile(fileName string){
	file,_ := os.Open(fileName)
	bufio := bufio.NewReader(file)
	bs,_,_ := bufio.ReadLine()
	array := make([]int,0,0)
	suffix :=1
	for len(bs)>0{
		if len(array) == NUMBER_TO_SORT {
			//排序
			array = QuickSort2(array)
			// 写入文件
			WriteTempFile(array,suffix)
			//重新创建数组
			array = make([]int,0)
			suffix++
		}
		number ,_ := strconv.Atoi(string(bs))
		array = append(array,number)
		bs,_,_ = bufio.ReadLine()
	}
	//对最后不足100数组排序，写入文件
	array = QuickSort2(array)
	WriteTempFile(array,suffix)
	//对各个小文件开始归并排序
	MergerSortFile("E://temp/")
	defer file.Close()
}
/**
	多个已经排好序的小文件合并排序
 */
func MergerSortFile(childFilePath string) {
	//最终生成的排序文件
	sortFile,_ := os.Create("E://number_sort.txt")
	sortFileBufio := bufio.NewWriter(sortFile)
	//过个排好序的小文件数组
	fileInfos,_ := ioutil.ReadDir(childFilePath)
	if len(fileInfos) == 0 {
		return
	}
	//用来存放每个小文件的最小值
	minArray := make([]int,len(fileInfos))
	fileArray := make([]*os.File,len(fileInfos))
	//每个小文件的bufio指针
	fileBufio := make([]*bufio.Reader,len(fileInfos))
	//用来标识小文件是否读到末尾
	isEnd := make([]bool,len(fileInfos))
	for index,fileInfo := range fileInfos{
		file,_ := os.Open("E://temp/"+fileInfo.Name())
		fileArray[index] = file
		fileBufio[index] = bufio.NewReader(file)
	}
	for{
		flag := true
		for i:=0;i<len(fileArray);i++{
			if (!isEnd[i] && minArray[i]==0){
				bs,_,_ := fileBufio[i].ReadLine()
				if(len(bs) == 0){
					isEnd[i] = true
				}
				value,_:=strconv.Atoi(string(bs))
				minArray[i] = value
			}
		}
		//找出多个小文件的最小值的最小值
		minValue := GetMin(minArray)
		//写入最终文件中
		if minValue != 0{
			sortFileBufio.WriteString(strconv.Itoa(minValue)+"\r\n")
			sortFileBufio.Flush()
		}
		//判断所有文件是否读取完毕
		for _,temp := range isEnd {
			if !temp{
				flag = false
				break
			}
		}
		//如果完毕，终止循环
		if flag{
			break
		}
	}
	//关闭流
	defer func() {
		sortFile.Close()
		for _,f := range fileArray{
			f.Close()
		}
	}()
}
func GetMin(array []int) int {
	index := 0
	min:=array[0]
	for k,value := range  array{
		if min > value && value != 0{
			min =  value
			index = k
		}
	}
	array[index] = 0
	return min
}
func WriteTempFile(array []int, i int) {
	os.MkdirAll(path.Dir("E://temp/"),os.ModePerm)
	file,_ := os.Create("E://temp/temp"+strconv.Itoa(i)+".txt")
	bufio := bufio.NewWriter(file)
	for _,v:= range array{
		bufio.WriteString(strconv.Itoa(v)+"\r\n")
	}
	bufio.Flush()
	defer file.Close()
}
func QuickSort2(array []int)[]int{
	if len(array) <= 1 {
		return array
	}
	middle,index := array[0],1
	head,tail := 0,len(array)-1
	for head<tail{
		if array[index] > middle{
			array[index],array[tail] = array[tail],array[index]
			tail--
		}else{
			array[index],array[head] = array[head],array[index]
			index++
			head++
		}
	}
	QuickSort2(array[:head])
	QuickSort2(array[head+1:])
	return array
}

func CreateSourceFile()  {
	file,_ := os.Create("E://number.txt")
	bufio := bufio.NewWriter(file)
	i:=10000
	for i>0{
		bufio.WriteString(strconv.Itoa(i)+"\r\n")
		i--
	}
	bufio.Flush()
	defer file.Close()
}