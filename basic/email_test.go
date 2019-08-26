package basic

import (
	"testing"
)

func Test_SendMail(t *testing.T) {

	SendMail(&Email{"sunpengwei1992@aliyun.com", "golang测试邮件", "朋伟：您好 \n \t这是golang的测试邮件"})

}
