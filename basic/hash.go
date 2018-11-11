package basic

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

/**
golang hash函数学习
*/

// 生成32位MD5
func Md5() {
	str := "sun_peng_wei"
	//方法一
	ctx := md5.New()
	ctx.Write([]byte(str))
	fmt.Println(hex.EncodeToString(ctx.Sum(nil)))
	//方法二
	has := md5.Sum([]byte(str))
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println(md5str1)
}
