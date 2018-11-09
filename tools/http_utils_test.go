package tools

import (
	"fmt"
	"go_common/excel"
	"strconv"
	"testing"
)

func TestHttpGetFile(t *testing.T) {

	for i:=2;i<=9;i++{
		if i != 4{
			data,_ := excel.ReadExcel("E:\\"+strconv.Itoa(i)+".xlsx")
			for _,sheet := range data{
				for _,row := range sheet{
					for celIndex,cel := range row{
						if celIndex == 1{
							if cel ==""{
								fmt.Println(i,"月份 ",row[0]," jss key 为空")
							}else{
								HttpGetFile(""+cel,"E:\\oss\\"+strconv.Itoa(i)+"\\")
							}
						}
					}
				}
			}
		}
	}
}

