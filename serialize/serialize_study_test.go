package serialize

import (
	"fmt"
	"testing"
	"time"
)

//987
func TestJsonSerialize(t *testing.T) {
	start := time.Now().UnixNano()
	bs := JsonSerialize()
	end := time.Now().UnixNano()
	fmt.Println((end - start), "ns", len(bs))
	fmt.Println(string(bs[1:7]))
}

func TestJsonDeSerialize(t *testing.T) {
	start := time.Now().UnixNano()
	JsonDeSerialize(jsonBs)
	end := time.Now().UnixNano()
	fmt.Println((end - start), "ns")

}

//396
func TestProto3Serialize(t *testing.T) {
	start := time.Now().UnixNano()
	bs := Proto3Serialize()
	end := time.Now().UnixNano()
	fmt.Println((end - start), "ns", len(bs))
}

//750
func TestMsgPackSerialize(t *testing.T) {
	start := time.Now().UnixNano()
	bs := MsgPackSerialize()
	end := time.Now().UnixNano()
	fmt.Println((end - start), "ns", len(bs), bs)
}

//1287
func TestXmlSerialize(t *testing.T) {
	start := time.Now().UnixNano()
	bs := XmlSerialize()
	end := time.Now().UnixNano()
	fmt.Println((end - start), "ns", len(bs))
	fmt.Println(string(bs))
}

func BenchmarkProto3Serialize(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Proto3Serialize()
	}
}

var protoBs = Proto3Serialize()

func BenchmarkProto3DeSerialize(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Proto3DeSerialize(protoBs)
	}
}

func BenchmarkJsonSerialize(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JsonSerialize()
	}
}

var jsonBs = JsonSerialize()

func BenchmarkJsonDeSerialize(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JsonDeSerialize(jsonBs)
	}
}

func BenchmarkMsgPackSerialize(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MsgPackSerialize()
	}
}

var msgPackJson = MsgPackSerialize()

func BenchmarkMsgPackDeSerialize(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MsgPackDeSerialize(msgPackJson)
	}
}

func BenchmarkXmlSerialize(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		XmlSerialize()
	}
}

var xmlBs = XmlSerialize()

func BenchmarkXmlDeSerialize(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		XmlDeSerialize(xmlBs)
	}
}
