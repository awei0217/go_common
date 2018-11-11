package algorithm

import (
	"fmt"
	"testing"
)

func TestNewBloomFilter(t *testing.T) {
	bloomFilter := NewBloomFilter()
	bloomFilter.add("123")
	fmt.Println(bloomFilter.contains("123"))
}
