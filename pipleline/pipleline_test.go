package pipleline

import (
	"fmt"
	"log"
	"testing"
)

func TestPipeline_Async(t *testing.T) {

	//工序(1)在pipeline外执行，最后一个工序是保存checkpoint
	pipeline := NewPipeline(8, 32, 2, 1)
	for {
		i := 1
		//(1)
		//加载100条数据，并修改变量checkpoint
		//data是数组，每个元素是一条评论，之后的联表、NLP都直接修改data里的每条记录。

		ok := pipeline.Async(func() error {
			i = i+1
			return nil
		}, func() error {
			i = i+1
			return nil
		}, func() error {
			i = i+1
			return nil
		}, func() error {
			fmt.Println(i)
			return nil
		})
		if !ok {
			break
		}
	}
	err := pipeline.Wait()
	if err != nil { log.Print(err) }
}
