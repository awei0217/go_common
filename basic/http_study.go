package basic

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

func HttpGet() {

	res, err := http.Get("http://www.baidu.com")
	checkErr(err)
	//将http请求相应的内容复制到标准输出（控制台）
	io.Copy(os.Stdout, res.Body)
	defer res.Body.Close()
}

func HttpPort() {
	res, err := http.Post("http://www.baidu.com",
		"application/x-www-form-urlencoded", strings.NewReader("name=spw&pwd=123"))
	checkErr(err)
	io.Copy(os.Stdout, res.Body)
	defer res.Body.Close()
}

func HttpPortForm() {
	res, err := http.PostForm("http://wwww.baidu.com", url.Values{"name": {"spw"}, "pwd": {"123"}})
	checkErr(err)
	io.Copy(os.Stdout, res.Body)
	defer res.Body.Close()
}

func HttpDo() {
	client := &http.Client{}
	jar, err := cookiejar.New(nil)
	checkErr(err)
	request, err := http.NewRequest("port", "http://www.baidu.com", strings.NewReader("name=spw&pwd=123"))
	checkErr(err)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client.Jar = jar
	resp, err := client.Do(request)
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
	defer resp.Body.Close()

}

func checkErr(err error) {
	if err != nil {
		log.Fatal("http err:", err)
	}
}
