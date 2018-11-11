package excel

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func TestReadExcel(t *testing.T) {
	sql := "insert into io_1_201809 " +
		"(status, exp_date, cky_no, store_no, action_code, record_type, job_no, exp_no, goods_no, qty, " +
		" action_time, service_type, region_no, data_source, yn, create_time) values"
	var buf bytes.Buffer
	data, err := ReadExcel("E:\\test\\10166_2018-06-01-2018-06-30_BillingDetail1.xls")
	if err != nil {
		fmt.Println("excel读取失败", err)
		return
	}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			for m := 0; m < len(data[i][j]); m++ {
				if m == 0 {
					buf.WriteString("(0," + "'" + data[i][j][m] + "'" + ",")
				} else if m == 9 {
					buf.WriteString("'" + data[i][j][m] + "',1,30,'bi_50_2',1,now()),")
				} else {
					buf.WriteString("'" + data[i][j][m] + "',")
				}
			}
		}
	}
	sql = sql + buf.String()
	t.Log(sql[0 : len(sql)-1])
}

func TestUnZipFile(t *testing.T) {
	for i := 2; i <= 9; i++ {
		if i != 4 {
			UnZipFile("E:\\oss\\"+strconv.Itoa(i)+"\\", "E://oss_oss//"+strconv.Itoa(i)+"//")
		}
	}
}
