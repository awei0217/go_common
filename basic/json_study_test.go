package basic

import (
	"fmt"
	"log"
	"testing"
	"time"
)

var stu = &Student{
	Id:       100001,
	Name:     "sunpengwei",
	Age:      26,
	Phone:    "18091772262",
	Score:    map[string]float32{"语文": 99.12, "数学": 100},
	Birthday: JsonTime(time.Now()),
}

func TestToJson(t *testing.T) {
	log.Println(ToJson(stu))
}

var stuESC = &Student{
	Id:       100001,
	Name:     "<sunpengwei>",
	Age:      26,
	Phone:    "18091772262",
	Score:    map[string]float32{"语文": 99.12, "数学": 100},
	Birthday: JsonTime(time.Now()),
}

func TestToJsonSkipESC(t *testing.T) {
	log.Println(ToJsonSkipESC(stuESC))
}

func TestFromJson(t *testing.T) {
	jsonStr, err := ToJson(stu)
	if err != nil {
		log.Fatalln(err)
	}
	result := &Student{}
	err = FromJson([]byte(jsonStr), result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}

func TestJsonMap(t *testing.T) {
	JsonMap()
}
