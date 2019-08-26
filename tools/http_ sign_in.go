package tools

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func NewClient() *http.Client {
	CurCookieJar, _ := cookiejar.New(nil)
	return &http.Client{
		Jar: CurCookieJar,
	}
}

//登录结果实体
type LoginResponse struct {
	Code int      `json:"code"`
	Data AuthInfo `json:"data"`
}

type AuthInfo struct {
	AuthToken string `json:"authToken"`
}

/**
client http请求客户端
authToken 请求头中必须携带的token
error 如果登陆失败，返回的错误信息
*/
func Login(param map[string]string, loginUrl string) (*http.Client, string, error) {

	values := url.Values{}
	for k, v := range param {
		values.Add(k, v)
	}
	req, err := http.NewRequest("POST", loginUrl, strings.NewReader(values.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := NewClient()
	// 执行登录操作
	res, err := client.Do(req)
	if nil != err {
		log.Fatal(err)
	}
	defer res.Body.Close()

	//登录返回结果
	bs, _ := ioutil.ReadAll(res.Body)
	result := string(bs)

	//对结果验证
	if !strings.Contains(result, "操作成功") {
		return nil, "", errors.New(result)
	}
	//将结果信息反序列化
	loginResponse := &LoginResponse{}
	err = json.Unmarshal(bs, loginResponse)
	if err != nil {
		return nil, "", errors.New("登录结果信息:" + result + "反序列化LoginResponse错误,错误信息:" + err.Error())
	}
	//返回client，authToken
	return client, loginResponse.Data.AuthToken, nil
}
