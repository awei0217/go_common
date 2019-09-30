package tools

import (
	"fmt"
	"go_common/excel"
	"sync"
	"testing"
)

func TestHttpGetFile(t *testing.T) {

	data, _ := excel.ReadExcel("E://spw.xlsx")
	for _, sheet := range data {
		for _, row := range sheet {
			for celIndex, cel := range row {
				if celIndex == 0 {
					if cel == "" {
						fmt.Println(row[0])
					} else {
						fmt.Println(cel)
						HttpGetFile(""+cel, "E:\\oss\\")
					}
				}
			}
		}
	}
}

func BenchmarkPostRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go func() {
			url := "http://127.0.0.1:8081/api/calculation/theaReplenishment"
			PostRequest(url)
		}()
	}
}

func TestPostRequest(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			url := "http://10.112.6.137/cs/api/calculation/theaReplenishment"
			PostRequest(url)
			wg.Done()
		}()
	}
	wg.Wait()
}
