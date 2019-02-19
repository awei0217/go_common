package basic

import (
	"fmt"
	"math/rand"
	"time"
)

func Rand(){

	//设置一个随机数种子
	rand.Seed(time.Now().Unix())


	//返回一个非负的伪随机int值。
	fmt.Println(rand.Int())
	//返回一个int32类型的非负的31位伪随机数。
	fmt.Println(rand.Int31())
	//返回一个int64类型的非负的63位伪随机数。
	fmt.Println(rand.Int63())




	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Int31n(100))
	fmt.Println(rand.Int63n(100))

	//返回一个0.0 到1.0的float32随机数
	fmt.Println(rand.Float32())
	//返回一个0.0 到1.0的float64随机数
	fmt.Println(rand.Float64())


	//返回一个[0,10) 范围的10个随机数
	fmt.Println(rand.Perm(10))

}
