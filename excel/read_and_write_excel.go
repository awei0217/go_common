package excel

import (
	"github.com/Luxurioust/excelize"
)

/**
	返回的数据三维切片，第一维的大小是 sheet  第二维的大小是rows,不包含表头,第三维的大小是cells,一行多少列
 */
func ReadExcel(filePath string)([][][]string,error ) {
	excel, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil,err
	}
	sheetMap := excel.GetSheetMap()
	sheets := make([][][]string,len(sheetMap))
	for index,value :=range  sheetMap{
		rows := excel.GetRows(value)
		rowsString := make([][]string,len(rows)-1)
		for key, row := range rows {
			if key == 0 { // 第一行是表头
				continue
			}
			//循环工作表行数的每一列
			cellsString := make([]string,len(row))
			for k, cell := range row {
				cellsString[k] = cell
			}
			rowsString[key-1] = cellsString
		}
		sheets[index-1] = rowsString
	}
	return sheets,nil
}


