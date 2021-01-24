package basic

import (
	"bytes"
	"encoding/json"
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"log"
	"time"
)

type JsonTime time.Time

const timeFormart = "2006-01-02 15:04:05"

func (t JsonTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

//UnmarshalJSON
func (t *JsonTime) UnmarshalJSON(bs []byte) error {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(bs), time.Local)
	if err != nil {
		return err
	}
	*t = JsonTime(now)
	return nil
}

//重写 string方法，方便打印调试
func (t JsonTime) String() string {
	return time.Time(t).Format(timeFormart)
}

type Person struct {
	StudentOne StudentOne `json:"studentOne"`
}
type StudentOne struct {
	Name string `json:"name"`
}

func JsonMap() {
	params := make(map[string]interface{})
	params["person"] = Person{StudentOne: StudentOne{Name: "ss"}}
	data, err := json.Marshal(params)

	fmt.Println(string(data), err)

	temp := make(map[string]interface{})

	json.Unmarshal(data, &temp)
	fmt.Println(temp)
}

//omitempty  这个属性可以在序列化时忽略0值和nil值
// score 不能序列化，属于私有字段
type Student struct {
	Id       int                `json:"id"` //如果想json后的id是字符串可以这样写  `json:"id,string"`
	Name     string             `json:"name"`
	Age      int                `json:"age"`
	Score    map[string]float32 `json:"score"`
	Phone    string             `json:"phone,omitempty"`
	Birthday JsonTime           `json:"birthday"`
}

func ToJson(v interface{}) (string, error) {
	result, err := json.Marshal(v)
	if err != nil {
		return "", nil
	}
	return string(result), nil
}

//跳过转义字符
func ToJsonSkipESC(v interface{}) (string, error) {
	bf := bytes.NewBuffer([]byte{})
	encode := json.NewEncoder(bf)
	encode.SetEscapeHTML(false)
	err := encode.Encode(v)
	if err != nil {
		return "", err
	}
	return bf.String(), nil
}

func FromJson(jsonStr []byte, v interface{}) error {
	err := json.Unmarshal(jsonStr, v)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func JsonGetValueByKey() {
	s := `{
		"tagA" : "json string",
		"tagB" : 1024,
		"tagD" : {
			"tagE":1000
		},
		"tagF":[
			"json array",
			1024,
			{"tagH":"json object"}
		]
	}`

	var i interface{}
	err := json.Unmarshal([]byte(s), &i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	AssertMap(i)
}

func AssertMap(i interface{}) {
	switch t := i.(type) {
	case map[string]interface{}:
		for k, v := range t {
			switch t1 := v.(type) {
			case map[string]interface{}:
				AssertMap(t1)
			case []interface{}:
				for k1, v1 := range t1 {
					switch t2 := v1.(type) {
					case map[string]interface{}:
						AssertMap(t2)
					default:
						fmt.Println(k1, ":", v1)
					}
				}
			default:
				fmt.Println(k, ":", v)
			}
		}
	}
}

func SimpleJsonStudy() {
	s := `{
		"tagA" : "json string",
		"tagB" : 1024,
		"tagD" : {
			"tagE":1000
		},
		"tagF":[
			"json array",
			1024,
			{"tagH":"json object"}
		]
	}`

	res, err := simplejson.NewJson([]byte(s))

	if err != nil {
		fmt.Println(err)
	}

	j, err := res.Get("tagF").GetIndex(2).Get("tagH").String()
	fmt.Println(j)

	k, err := res.GetPath("tagD", "tagE").Int64()
	fmt.Println(k)

	if l, ok := res.CheckGet("tagD"); ok { //true ,有该字段
		fmt.Println(ok, l)
	}
}
