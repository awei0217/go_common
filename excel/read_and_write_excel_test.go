package excel

import (
	"testing"
)

func TestReadExcel(t *testing.T) {
	data := ReadExcel("E:\\111.xlsx")
	for i:=0;i<len(data);i++{
		for j:=0;j<len(data[i]);j++{
			t.Log(len(data[i][j]))
		}
	}
}
