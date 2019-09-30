package tools

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	param := make(map[string]string)

	param["userName"] = "32c3VucGVuZ3dlaTE=111"
	param["passWord"] = "32QWJjZDEyMw==222"
	param["domainCode"] = "GOMEDQ"
	client, authToken, err := Login(param, "http://office.gome.com.cn/newWeb/auth/domainsLogin.do")
	if err != nil {
		t.Error(err)
	}

	//利用登录成功返回的client继续发送请求
	req, _ := http.NewRequest("GET", "http://office.gome.com.cn/newWeb/data/gateway/site/getuserinfo.do?userId=1181345", nil)
	//添加请求认证信息
	req.Header.Add("Authorization", authToken)
	res, err := client.Do(req)
	defer res.Body.Close()
	bs, _ := ioutil.ReadAll(res.Body)
	result := string(bs)
	t.Log(result)
}
