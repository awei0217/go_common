package tools

import (
	"fmt"
	"go_common/excel"
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
