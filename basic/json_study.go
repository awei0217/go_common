package basic

import (
	"bytes"
	"encoding/json"
	"log"
	"time"
)



type JsonTime time.Time

const timeFormart  =  "2006-01-02 15:04:05"

func (t JsonTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}
//UnmarshalJSON
func (t *JsonTime) UnmarshalJSON(bs []byte)error{
	now,err := time.ParseInLocation(`"`+timeFormart+`"`,string(bs),time.Local)
	if err != nil{
		return err
	}
	*t = JsonTime(now)
	return nil
}
//重写 string方法，方便打印调试
func (t JsonTime)String()string{
	return time.Time(t).Format(timeFormart)
}


//omitempty  这个属性可以在序列化时忽略0值和nil值
// score 不能序列化，属于私有字段
type Student struct {
	Id int `json:"id"` //如果想json后的id是字符串可以这样写  `json:"id,string"`
	Name string `json:"name"`
	Age int `json:"age"`
	Score map[string]float32 `json:"score"`
	Phone string `json:"phone,omitempty"`
	Birthday JsonTime `json:"birthday"`
}

func ToJson(v interface{})(string,error){
	result,err := json.Marshal(v)
	if err != nil{
		return "",nil
	}
	return string(result),nil
}

//跳过转义字符
func ToJsonSkipESC(v interface{})(string,error){
	bf := bytes.NewBuffer([]byte{})
	encode := json.NewEncoder(bf)
	encode.SetEscapeHTML(false)
	err := encode.Encode(v)
	if err != nil{
		return "",err
	}
	return bf.String(),nil
}



func FromJson (jsonStr []byte,v interface{})error{
	err := json.Unmarshal(jsonStr,v)
	if err != nil{
		log.Fatalln(err)
		return err
	}
	return nil
}

