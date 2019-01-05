package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var conn redis.Conn

//redis文章
//https://www.cnblogs.com/276815076/p/7245333.html

/**
初始化
*/
func init() {
	conn, _ = redis.Dial("tcp4", "127.0.0.1:6369", redis.DialPassword(""))
}
func String() {
	// set key value 原子性操作
	conn.Send("set", "spw", "spw")
	// exists key  存在返回1 不存在返回0
	re0, _ := conn.Do("exists", "spw")
	fmt.Println(re0.(interface{}))
	re1, _ := conn.Do("get", "spw")
	fmt.Println(string(re1.([]byte)))

	//mset key1 value1 key2 value2  是一个原子性操作，如果key已经存在，则覆盖其值
	conn.Send("mset", "spw1", "spw1", "spw2", "spw2")
	re2, _ := conn.Do("mget", "spw1", "spw2")
	fmt.Println(string(re2.([]interface{})[0].([]byte)))
	fmt.Println(string(re2.([]interface{})[1].([]byte)))

	// setex key second value 设置key的值为value 指定second秒过期
	conn.Send("SETEX", "spw3", 0.1, "spw3")
	time.Sleep(time.Second * 1)
	re3, _ := conn.Do("get", "spw3")
	fmt.Println(re3)

	//getset key value 将给定key 设置value值，并返回旧值
	re4, _ := conn.Do("getset", "spw1", "spw11")
	fmt.Println(string(re4.([]byte)))

	defer conn.Close()
}

func Hash() {
	conn.Send("hset", "student1", "name", "spw")
	conn.Send("hset", "student1", "age", 25)
	conn.Send("hset", "student1", "sex", "nan")

	re1, _ := conn.Do("hget", "student1", "name")
	re2, _ := conn.Do("hkeys", "student1")
	re3, _ := conn.Do("hgetall", "student1")
	re4, _ := conn.Do("hexists", "student1", "name")
	fmt.Println(string(re1.([]byte)))
	fmt.Println(re2.([]interface{}))
	fmt.Println(re3.([]interface{}))
	fmt.Println(re4.(interface{}))
	defer conn.Close()
}

func List() {
	//将一个或多个值 value 插入到列表 key 的表头,当 key 存在但不是列表类型时，返回一个错误
	conn.Do("lpush", "spw10", "1")
	conn.Do("lpush", "spw10", "1")
	conn.Do("lpush", "spw10", "2")
	conn.Do("lpush", "spw10", "3", "4", "5")
	//移除并返回列表 key 的头元素。 当 key 不存在时，返回 nil
	re1, _ := conn.Do("lpop", "spw10")
	//移除并返回列表 key 的尾元素。当 key 不存在时，返回 nil
	re2, _ := conn.Do("rpop", "spw10")
	fmt.Println(re1.(interface{}))
	fmt.Println(re2.(interface{}))

	//将值 value 插入到列表 key 的表头，当且仅当 key 存在并且是一个列表。//和 LPUSH 命令相反，当 key 不存在时， LPUSHX 命令什么也不做。
	conn.Do("lpushx", "spw10", "20")
	//返回列表的长度
	re3, _ := conn.Do("llen", "spw10")
	fmt.Println(re3.(interface{}).(int64))

	//将列表 key 下标为 index 的元素的值设置为 value 。当 index 参数超出范围，或对一个空列表( key 不存在)进行 LSET 时，返回一个错误。
	conn.Send("lset", 0, "100")

	//阻塞方法 当给定列表内没有任何元素可供弹出的时候，连接将被 BLPOP 命令阻塞，直到等待超时或发现可弹出元素为止
	//超时参数 timeout 接受一个以秒为单位的数字作为值。超时参数设为 0 表示阻塞时间可以无限期延长(block indefinitely) 。
	re4, _ := conn.Do("BLPOP", "spw10", 0) //jimdb 貌似不支持
	fmt.Println(re4)

	defer conn.Close()
}

func Set() {
	// 将一个或多个member 元素加入到集合 key 当中，已经存在于集合的 member 元素将被忽略。
	conn.Send("sadd", "key2", "value2")
	conn.Send("sadd", "key2", "value3")
	//返回集合 key 的基数(集合中元素的数量)
	re1, _ := conn.Do("scard", "key2")
	fmt.Println(re1.(interface{}))
	//返回集合 key 中的所有成员
	re2, _ := conn.Do("smembers", "key2")
	fmt.Println(re2.(interface{}))
	//spop 移除并返回集合中的一个随机元素
	re3, _ := conn.Do("spop", "key2")
	fmt.Println(re3.(interface{}))
	//移除集合 key 中的一个或多个 member 元素，不存在的 member 元素会被忽略
	re4, _ := conn.Do("srem", "key2", "value3")
	fmt.Println(re4.(interface{}))
}

/**
HyperLogLog 可以接受多个元素作为输入，并给出输入元素的基数的估算值，
基数：集合中不同元素的数量。比如 {'apple', 'banana', 'cherry', 'banana', 'apple'} 的基数就是 3
基数计数(cardinality counting)通常用来统计一个集合中不重复的元素个数
估算值：算法给出的基数并不是精确的，可能会比实际稍微多一些或者稍微少一些，但会控制在合理的范围之内
HyperLogLog 的优点是 即使输入的元素的数量非常多，计算基数所需的空间总是固定的，并且很小，
每个 HyperLogLog 的键只需要花费 12 KB 内存，就可以计算接近 2^64 个不同元素的基数
但是，因为 HyperLogLog 只会根据输入元素来计算基数，而不会储存输入元素本身，所以HyperLogLog 不能像集合那样，返回输入的各个元素
HyperLogLog 原理  https://blog.csdn.net/firenet1/article/details/77247649
*/
func HyperLogLog() {
	//将元素添加至hyperLogLog,pfadd这个命令会对hyperloglog进行修改，如果添加元素后，基数的估算值发生了变化，返回1，否则返回0
	re1, _ := conn.Do("pfadd", "hyperloglog_key", 0)
	re2, _ := conn.Do("pfadd", "hyperloglog_key", 0)
	fmt.Println(re1.(interface{})) // 1
	fmt.Println(re2.(interface{})) // 0

	for i := 0; i < 10; i++ {
		conn.Do("pfadd", "hyperloglog_key", i)
	}
	//pfcount 返回HyperLogLog 给定key的基数估算值
	re4, _ := conn.Do("pfcount", "hyperloglog_key")
	fmt.Println(re4.(interface{}))
	// 当给定多个 HyperLogLog 时，命令会先对给定的 HyperLogLog 进行并集计算，得出一个合并后的HyperLogLog ，然后返回这个合并 HyperLogLog 的基数估算值作为命令的结果（合并得出的HyperLogLog 不会被储存，使用之后就会被删掉）
	re5, _ := conn.Do("pfmerge", "hyperloglog_key1", "hyperloglog_key2")
	fmt.Println(re5.(interface{}))
	re3, _ := conn.Do("expire", "hyperloglog_key", 1)
	fmt.Println(re3.(interface{}))
}

func SortedSet() {
	//将一个或多个 member 元素及其 score 值加入到有序集 key 当中 ,score 值可以是整数值或双精度浮点数。
	conn.Send("zadd", "zkey", 1, "z1")
	conn.Send("zadd", "zkey", 2, "z2")
	conn.Send("zadd", "zkey", 3, "z3")
	conn.Send("zadd", "zkey", -9, "z9")
	conn.Send("zadd", "zkey", 10, "z10")
	conn.Send("zadd", "zkey", 4, "z4")
	conn.Send("zadd", "zkey", 6, "z6")
	//返回有序集 key 的数量。当 key 不存在时，返回 0 。
	re1, _ := conn.Do("zcard", "zkey")
	fmt.Println(re1.(interface{}))
	//返回有序集 key 中， score 值在 min 和 max 之间(默认包括 score 值等于 min 或 max )的成员的数量
	re2, _ := conn.Do("zcount", "zkey", 1, 3)
	fmt.Println(re2.(interface{}))
	// 返回有序集中所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。有序集成员按 score 值递增(从小到大)次序排列
	//min 和 max 可以是 -inf 和 +inf  代表返回所有有序集
	//区间的取值使用闭区间 (小于等于或大于等于)，你也可以通过给参数前增加 ( 符号来使用可选的开区间 (小于或大于) ZRANGEBYSCORE zset (1 5  返回所有符合条件 1 < score <= 5 的成员
	re3, _ := conn.Do("zrangebyscore", "zkey", 1, 5)    // 1<=  score <= 5
	re4, _ := conn.Do("zrangebyscore", "zkey", "(1", 5) // 1 < score <= 5
	fmt.Println(re3.(interface{}))
	fmt.Println(re4.(interface{}))
	//zrangebyscore 获取键为zkey的 socre 在6,500 之间的value的第一条
	re5, _ := conn.Do("zrangebyscore", "zkey", 6, 500, "limit", 0, 1)
	fmt.Println(re5.(interface{}))
	//返回有序集 key 中，成员 member 的 score 值。
	re6, _ := conn.Do("zscore", "zkey", "z1")
	fmt.Println(re6.(interface{}))
}

/**
redis 位图
*/
func BitMap() {
	// setbit 将给定key某一位设置为0 或者 1，第一次设置时返回0，后面再对同一位设置时返回1 ; 最大长度为 2^32位数组
	re1, _ := conn.Do("setbit", "bitmapkey", 70, 1)
	fmt.Println(re1.(interface{}))
	//getbit 获取给定key的某一位上的值
	re2, _ := conn.Do("getbit", "bitmapkey", 60)
	fmt.Println(re2.(interface{}))
	//bitcount 用来统计指定范围内内1的个数
	re3, _ := conn.Do("bitcount", "bitmapkey", 0, 100)
	fmt.Println(re3.(interface{}))
	//bitops 用来查找指定范围内出现的第一个0 或者 1 ,第二个参数代表你要查的是0还是1 ，后面两个是范围
	re4, _ := conn.Do("bitpos", "bitmapkey", 0, 50, 100)
	fmt.Println(re4.(interface{}))
}

/**
https://www.cnblogs.com/KevinYang/archive/2009/02/01/1381803.html
redis 布隆过滤器 在redis4.0 版本中才存在
BloomFilter 对于已经存进去的元素不会误判，但对于没有存进去的元素可能存在误判
*/
func BloomFilter() {
	for i := 0; i < 1000; i++ {
		// 往布隆过滤器中指定的key添加元素
		re2, _ := conn.Do("bf.add", "bloom_key", i)
		fmt.Println(re2.(interface{}))
	}

	for i := 0; i < 10000; i++ {
		//判断布隆过滤器中指定key的value是否存在
		re1, _ := conn.Do("bf.exists", "bloom_key", i)
		fmt.Println(re1.(interface{}))
	}
}

/**
设置分布式锁的正确姿势
*/
func SetDistributeLock() {
	//key1 锁的key 1 锁的过期时间  requestId 设置锁的请求ID
	re1, _ := conn.Do("setex", "key1", 10, "requestId")
	fmt.Println(re1.(interface{}))
}

/**
释放分布式锁的正确姿势
*/
func ReleaseDistributeLock() {
	script := "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end"
	// 1 代表key的个数
	re1, _ := redis.NewScript(1, script).Do(conn, "key1", "requestId")
	fmt.Println(re1.(interface{}))
	re2, _ := conn.Do("get", "key1")
	fmt.Println(re2)
}

/**
redis实现延时队列
*/
func DelayedQueue() {
	//添加一个任务到队列中,score值为当前时间秒 + 10 ，代表延迟10秒执行
	conn.Do("zadd", "zkey1", time.Now().Unix()+10, "z1")
	//添加一个任务到队列中,score值为当前时间秒 + 20 ，代表延迟20秒执行
	conn.Do("zadd", "zkey1", time.Now().Unix()+20, "z2")
	// 去除0-当前表的有序集元素
	for {
		// 每次只取一条
		re1, _ := conn.Do("zrangebyscore", "zkey1", 0, time.Now().Unix(), "limit", 0, 1)
		if re1 == nil || len(re1.([]interface{})) == 0 {
			// 如果空 ，休息一秒钟,防止循环耗高CPU 和 调用redis频率过快，导致redis QPS过高
			time.Sleep(time.Second * 1)
		} else {
			//删除已经取出的任务 (核心，多线程可以竞争，成功的那个可以处理任务)
			re2, _ := conn.Do("zrem", "zkey1", string(re1.([]interface{})[0].([]byte)))
			if re2.(int64) == 1 {
				fmt.Println("模拟处理任务业务")
			}
		}
	}
}
