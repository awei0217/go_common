package excel

import (
	"github.com/Luxurioust/excelize"
	"log"
	"strconv"
)


func ReadExcel(filePath string)  {
	excel, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Println("ERR:"+err.Error())
	}
	index := excel.GetSheetIndex("sheet1")
	// Get all the rows in a sheet.
	rows := excel.GetRows("sheet" + strconv.Itoa(index))
	for key, row := range rows {
		if(key == 0){
			continue
		}
		//循环工作表行数的每一列
		var wsMap map[string]interface{}
		for k, cell := range row {
			log.Println(wsMap,k,cell)
		}
	}
	excel.SaveAs(filePath)

}
