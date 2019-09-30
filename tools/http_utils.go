package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func PostRequest(httpUrl string) {

	params := make(map[string]interface{})
	params["createTime"] = "2019-09-09 12:12:12"
	params["salegrpCodes"] = []string{"100"}
	params["modelCode"] = 1
	params["dealCode"] = 1
	params["skuCodeArray"] = []string{"000000000100461042"}
	params["businessUnitCode"] = "舒适家事业部"
	params["saleChannelType"] = "1"
	params["buyorg"] = 103
	params["purchaseType"] = "0001"

	data, err := json.Marshal(params)
	if err != nil {
		log.Println("序列化错误", err)
		return
	}
	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", httpUrl, body)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := NewClient()
	// 执行登录操作
	res, err := client.Do(req)
	if nil != err {
		log.Println("请求错误", err)
	}
	defer res.Body.Close()
}
