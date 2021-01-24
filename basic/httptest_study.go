package basic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ADDRESS = "shanghai"
)

type PersonNew struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
}

func GetInfo(api string) ([]PersonNew, error) {
	url := fmt.Sprintf("%s/person?addr=%s", api, ADDRESS)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return []PersonNew{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return []PersonNew{}, fmt.Errorf("get info didn’t respond 200 OK: %s", resp.Status)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	personList := make([]PersonNew, 0)
	err = json.Unmarshal(bodyBytes, &personList)
	if err != nil {
		return []PersonNew{}, fmt.Errorf("decode data fail")
	}
	return personList, nil
}

//解释一下：
//
//>我们通过httptest.NewServer创建了一个测试的http server
//
//>读请求设置通过变量r *http.Request，写变量（也就是返回值）通过w http.ResponseWriter
//
//>通过ts.URL来获取请求的URL（一般都是<http://ip:port>）
//
//>通过r.Method来获取请求的方法，来测试判断我们的请求方法是否正确
//
//>获取请求路径：r.URL.EscapedPath()，本例中的请求路径就是"/person"
//
//>获取请求参数：r.ParseForm，r.Form.Get("addr")
//
//>设置返回的状态码：w.WriteHeader(http.StatusOK)
//
//>设置返回的内容（这就是我们想要的结果）：w.Write(personResponseBytes)，注意w.Write()接收的参数是[]byte，因此需要将object对象列表通过json.Marshal(personResponse)转换成字节。
