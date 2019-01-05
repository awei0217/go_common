package excel

import (
	"archive/zip"
	"fmt"
	"github.com/Luxurioust/excelize"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

/**
返回的数据三维切片，第一维的大小是 sheet  第二维的大小是rows,不包含表头,第三维的大小是cells,一行多少列
*/
func ReadExcel(filePath string) ([][][]string, error) {
	excel, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	sheetMap := excel.GetSheetMap()
	sheets := make([][][]string, len(sheetMap))
	for index, value := range sheetMap {
		rows := excel.GetRows(value)
		if(len(rows)==0){
			fmt.Println(index,"sheet行数为空")
			continue
		}
		rowsString := make([][]string, len(rows)-1)
		for key, row := range rows {
			if key == 0 { // 第一行是表头
				continue
			}
			//循环工作表行数的每一列
			cellsString := make([]string, len(row))
			for k, cell := range row {
				cellsString[k] = cell
			}
			rowsString[key-1] = cellsString
		}
		sheets[index-1] = rowsString
	}
	return sheets, nil
}

/**
第一个参数的源目录
第二个参数是解压后存放解压文件的目录
*/
func UnZipFile(sourceDir string, targetDir string) {
	//获取这个目录下的所有压缩文件
	filePathArray := getAllFile(sourceDir)
	for m, fileName := range filePathArray {
		// m 是为了防止不同压缩文件中的文件名重复给文件名加编号
		_, err := unzip(fileName, targetDir, m+1)
		if err != nil {
			fmt.Println(fileName, "解压失败", err)
			continue
		}
	}
}
func getAllFile(pathname string) []string {
	filePath := make([]string, 0, 0)
	rd, _ := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			getAllFile(pathname + fi.Name() + "\\")
		} else {
			filePath = append(filePath, pathname+fi.Name())
		}
	}
	return filePath
}
func unzip(archive, target string, m int) (string, error) {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		fmt.Println(archive, "解析失败")
		return "", err
	}
	if err := os.MkdirAll(target, 0755); err != nil {
		return "", err
	}
	for _, file := range reader.File {
		path := filepath.Join(target, strconv.Itoa(m)+"_"+file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}
		fileReader, err := file.Open()
		if err != nil {
			return "", err
		}
		defer fileReader.Close()
		targetFile, err := os.Create(path)
		if err != nil {
			return "", err
		}
		defer targetFile.Close()
		if _, err := io.Copy(targetFile, fileReader); err == nil {
			return targetFile.Name(), nil
		}
	}
	return "", err
}
