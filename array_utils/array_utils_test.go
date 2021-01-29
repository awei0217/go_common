package array_utils

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("初始化资源")
	result := m.Run()
	fmt.Println("释放资源")
	os.Exit(result)
}
func TestFindMaxSeqSum(t *testing.T) {
	t.Cleanup(func() {
		fmt.Println("clear resource")
	})
	defer func() {
		fmt.Println("defer ")
	}()
	sum := FindMaxSeqSum([]int{1, 3, -9, 6, 8, -19})
	if sum == 14 {
		t.Log("successful")
	} else {
		t.Error("failed")
	}
}
func TestCleanup(t *testing.T) {
	defer func() { t.Log("defer resource1") }()
	t.Cleanup(func() {
		t.Log("clear resource")
	})
	defer func() { t.Log("defer resource2") }()
	t.Log("test cleanup")
}
func BenchmarkFindMaxSeqSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindMaxSeqSum([]int{1, 3, -9, 6, 8, -19})
	}
}

//func ExampleFindMaxSeqSum() {
//	FindMaxSeqSum([]int{1, 3, -9, 6, 8, -19})
//	// OutPut: 14
//}
