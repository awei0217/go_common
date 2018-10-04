package algorithm

/**
用数组实现LRU(最近最少使用)算法
*/
const DATA_SIZE = 10

type LRUArray struct {
	//存储缓存的key
	keys []string
	//存储缓存的value
	values []string
	//下标值
	maxIndex int
}

func (lru *LRUArray) Put(key, value string) {
	if lru.keys == nil || lru.values == nil {
		lru.keys = make([]string, DATA_SIZE)
		lru.values = make([]string, DATA_SIZE)
		lru.maxIndex = DATA_SIZE - 1
	}
	// 缓存满了，需要淘汰
	if lru.maxIndex < 0 {
		i := len(lru.keys) / 2
		l := len(lru.keys) - 1
		for j := i; j >= 0; j-- {
			lru.keys[l] = lru.keys[j]
			lru.values[l] = lru.values[j]
			l--
		}
		lru.maxIndex = len(lru.keys) - (i + 2)
	}
	lru.keys[lru.maxIndex] = key
	lru.values[lru.maxIndex] = value
	lru.maxIndex--
}

func (lru *LRUArray) Get(key string) string {
	for i := lru.maxIndex; i < len(lru.keys); i++ {
		if key == lru.keys[i] {
			v := lru.values[i]
			j := i
			for j > lru.maxIndex+1 {
				lru.keys[j] = lru.keys[j-1]
				lru.values[j] = lru.values[j-1]
				j--
			}
			lru.keys[j] = key
			lru.values[j] = v
			return v
		}
	}
	return ""
}
