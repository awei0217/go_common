package serialize

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/vmihailenco/msgpack"
)

type Person struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Sex      int       `json:"sex"`
	Like     []string  `json:"like"`
	Children []*Person `json:"children"`
	Address  []string  `json:"address"`
	Phone    []string  `json:"phone"`
	Card     string    `json:"card"`
	QQ       string    `json:"qq"`
	WeChat   string    `json:"we_chat"`
	Birthday time.Time `json:"birthday"`
	Money    float64   `json:"money"`
}

//正常的结构体
var p Person = Person{
	Name: "孙朋伟",
	Age:  27,
	Sex:  1,
	Like: []string{"Coding", "跑步", "读书", "旅行"},
	Children: []*Person{{Name: "孩子1", Age: 0, Sex: 0, Like: []string{"Coding", "跑步", "读书", "旅行"}, Children: nil, Address: []string{"西安市"}, Phone: []string{"188888888888"}, Birthday: time.Now()},
		{Name: "孩子1", Age: 0, Sex: 0, Like: []string{"Coding", "跑步", "读书", "旅行"}, Children: nil, Address: []string{"西安市"}, Phone: []string{"188888888888"}, Birthday: time.Now()},
		{Name: "孩子1", Age: 0, Sex: 0, Like: []string{"Coding", "跑步", "读书", "旅行"}, Children: nil, Address: []string{"西安市"}, Phone: []string{"188888888888"}, Birthday: time.Now()}},
	Address:  []string{"西安市", "北京市"},
	Phone:    []string{"188888888888"},
	Card:     "610525199202171330",
	QQ:       "1024886110",
	WeChat:   "1109220120",
	Birthday: time.Now(),
	Money:    987.134,
}

//proto3的结构体（和正常的结构体时间类型不一样，其它都一样）
var pp PersonProto3 = PersonProto3{
	Name: "孙朋伟",
	Age:  27,
	Sex:  1,
	Like: []string{"Coding", "跑步", "读书", "旅行"},
	Children: []*PersonProto3{{Name: "孩子1", Age: 0, Sex: 0, Like: []string{"Coding", "跑步", "读书", "旅行"}, Children: nil, Address: []string{"西安市"}, Phone: []string{"188888888888"}, Birthday: &timestamp.Timestamp{Seconds: time.Now().Unix(), Nanos: int32(time.Now().UnixNano())}},
		{Name: "孩子1", Age: 0, Sex: 0, Like: []string{"Coding", "跑步", "读书", "旅行"}, Children: nil, Address: []string{"西安市"}, Phone: []string{"188888888888"}, Birthday: &timestamp.Timestamp{Seconds: time.Now().Unix(), Nanos: int32(time.Now().UnixNano())}},
		{Name: "孩子1", Age: 0, Sex: 0, Like: []string{"Coding", "跑步", "读书", "旅行"}, Children: nil, Address: []string{"西安市"}, Phone: []string{"188888888888"}, Birthday: &timestamp.Timestamp{Seconds: time.Now().Unix(), Nanos: int32(time.Now().UnixNano())}}},
	Address:  []string{"西安市", "北京市"},
	Phone:    []string{"188888888888"},
	Card:     "610525199202171330",
	Qq:       "1024886110",
	WeChat:   "1109220120",
	Birthday: &timestamp.Timestamp{Seconds: time.Now().Unix(), Nanos: int32(time.Now().UnixNano())},
	Money:    987.134,
}

func JsonSerialize() []byte {
	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	return bs
}

func JsonDeSerialize(bs []byte) {
	p := &Person{}
	err := json.Unmarshal(bs, p)
	if err != nil {
		fmt.Println(err)
	}
}

func Proto3Serialize() []byte {
	bs, err := proto.Marshal(&pp)
	if err != nil {
		fmt.Println(err)
	}
	return bs
}

func Proto3DeSerialize(bs []byte) {
	m := &PersonProto3{}
	err := proto.Unmarshal(bs, m)
	if err != nil {
		fmt.Println(err)
	}
}

func MsgPackSerialize() []byte {
	bs, err := msgpack.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	//把数据写入文件
	ioutil.WriteFile("./person.txt", bs, os.ModePerm)
	return bs
}

func MsgPackDeSerialize(bs []byte) {
	m := &Person{}
	err := msgpack.Unmarshal(bs, m)
	if err != nil {
		fmt.Println(err)
	}
}

func XmlSerialize() []byte {
	bs := bytes.Buffer{}
	xml.NewEncoder(&bs).Encode(p)
	return bs.Bytes()
}

func XmlDeSerialize(bs []byte) {
	p := &Person{}
	err := xml.NewDecoder(bytes.NewBuffer(bs)).Decode(p)
	if err != nil {
		fmt.Println(err)
	}
}
