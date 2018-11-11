package tools

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

/**
使用httpUrl下载文件
第一个参数，http连接
第二参数，下载下来的文件存放路径
*/
func HttpGetFile(httpUrl string, filePath string) {
	if httpUrl == "" {
		return
	}
	uri, _ := url.ParseRequestURI(httpUrl)
	fileName := path.Base(uri.Path)
	res, err := http.Get(httpUrl)
	if err != nil {
		fmt.Println(httpUrl, "访问错误", err)
	}
	file, _ := os.Create(filePath + fileName)
	io.Copy(file, res.Body)
}

/**
  如果返回的错误为nil,说明文件或文件夹存在
  如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
  如果返回的错误为其它类型,则不确定是否在存在
*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
