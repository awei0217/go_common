package algorithm

import (
	"bytes"
)

//BitMap实现
type BitMap []uint64

const (
	Address_Bits_Per_Word uint8  = 6
	Words_Per_Size        uint64 = 64 //64位
)

/**
创建指定初始化大小的BitMap
*/
func NewBitMap(nbits int) *BitMap {
	//得出需要多少个unit64的长度，  一个unit64 占8个字节，一个字节占8位
	//右移6位相当于处于64
	wordsLen := (nbits - 1) >> Address_Bits_Per_Word
	temp := BitMap(make([]uint64, wordsLen+1, wordsLen+1))
	return &temp
}

/**
把指定位置设为ture
*/
func (this *BitMap) Set(bitIndex uint64) {
	wIndex := this.wordIndex(bitIndex)
	this.expandTo(wIndex)
	(*this)[wIndex] |= (uint64(0x01) << (bitIndex % Words_Per_Size))
}

//设置指定位置为false
func (this *BitMap) Clear(bitIndex uint64) {
	wIndex := this.wordIndex(bitIndex)
	if wIndex < len(*this) {
		(*this)[wIndex] &^= uint64(0x01) << (bitIndex % Words_Per_Size)
	}
}

//获取指定位置的值
func (this *BitMap) Get(bitIndex uint64) bool {
	wIndex := this.wordIndex(bitIndex)
	return (wIndex < len(*this)) && ((*this)[wIndex]&(uint64(0x01)<<(bitIndex%Words_Per_Size)) != 0)
}

/**
以二进制串的格式打印bitMap内容
*/
func (this *BitMap) ToString() string {
	var temp uint64
	strAppend := &bytes.Buffer{}
	for i := 0; i < len(*this); i++ {
		temp = (*this)[i]
		for j := 0; j < 64; j++ {
			if temp&(uint64(0x01)<<uint64(j)) != 0 {
				strAppend.WriteString("1")
			} else {
				strAppend.WriteString("0")
			}
		}
	}
	return strAppend.String()
}

//定位位置
func (this BitMap) wordIndex(bitIndex uint64) int {
	return int(bitIndex >> Address_Bits_Per_Word)
}

//扩容:每次扩容两倍
func (this *BitMap) expandTo(wordIndex int) {
	wordsRequired := wordIndex + 1
	if len(*this) < wordsRequired {
		if wordsRequired < 2*len(*this) {
			wordsRequired = 2 * len(*this)
		}
		newCap := make([]uint64, wordsRequired, wordsRequired)
		copy(newCap, *this)
		(*this) = newCap
	}
}
