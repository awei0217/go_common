package basic

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"hash/crc64"
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

func Base64() {
	//标准加密
	en := base64.StdEncoding.EncodeToString([]byte("sunpengwei"))
	//标准解密
	de, _ := base64.StdEncoding.DecodeString(en)
	fmt.Println(en, string(de))
}

func Hash() {
	// %x 16进制
	// %b 2进制
	// %s 字符串
	sha1 := sha1.Sum([]byte("sunpengwei"))
	fmt.Println(fmt.Sprintf("%x", sha1))

	sha2 := sha256.Sum256([]byte("sunpengwei"))
	fmt.Println(fmt.Sprintf("%x", sha2))

	sha3 := sha512.Sum512([]byte("sunpengwei"))
	fmt.Println(fmt.Sprintf("%x", sha3))

}

func Crc() {
	crc32_1 := crc32.Checksum([]byte("sunpengwei"), &crc32.Table{})
	crc32_2 := crc32.ChecksumIEEE([]byte("sunpengwei"))
	fmt.Println(crc32_1, crc32_2)

	crc64_1 := crc64.Checksum([]byte("sunpengwei"), &crc64.Table{})
	fmt.Println(crc64_1)
}
