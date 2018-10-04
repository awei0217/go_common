package algorithm

import (
	"hash/crc32"
)

type HashMap struct {
	loadFactor   float64
	rehashFactor int
	size         int
	data         []*Entry
}

type Entry struct {
	key   string
	value string
	next  *Entry
}

func NewHashMap() *HashMap {
	return &HashMap{
		loadFactor:   0.75,
		rehashFactor: 2,
		size:         0,
		data:         make([]*Entry, 16),
	}
}

/**
key 相同的 value会覆盖
*/
func (hm *HashMap) Put(key, value string) {
	if float64(hm.size/len(hm.data)) > hm.loadFactor {
		rehash(hm)
	}
	hash := hashCode(key)
	index := getIndex(hash, len(hm.data))
	entry := &Entry{
		key:   key,
		value: value,
	}
	oldEntry := hm.data[index]
	if oldEntry == nil {
		hm.data[index] = entry
	} else {

		for oldEntry != nil {
			if oldEntry.key == key {
				oldEntry.value = value
				return
			}
			if oldEntry.next == nil {
				oldEntry.next = entry
				break
			}
			oldEntry = oldEntry.next
		}
	}
	hm.size++
}

/**
如果第二个参数返回 true，代表key存在，并第一个参数返回value
如果第二个参数返回false 代表key不存在，并第一个参数返回空串
*/
func (hm *HashMap) Get(key string) (string, bool) {
	hash := hashCode(key)
	index := getIndex(hash, len(hm.data))
	en := hm.data[index]
	for en.next != nil {
		if en.key == key {
			return en.value, true
		}
		en = en.next
	}
	if en.key == key {
		return en.value, true
	}
	return "", false
}

func (hm *HashMap) Remove(key string) bool {

	return false
}

/**
hash 扩容
*/
func rehash(hm *HashMap) {
	newData := make([]*Entry, hm.rehashFactor*len(hm.data))
	for _, v := range hm.data {
		if v == nil {
			continue
		}
		if v != nil && v.next == nil {
			hash := hashCode(v.key)
			index := getIndex(hash, len(newData))
			newData[index] = v
			continue
		}
		for v != nil {
			hash := hashCode(v.key)
			index := getIndex(hash, len(newData))
			oldEntry := newData[index]
			if oldEntry == nil {
				en := &Entry{
					key:   v.key,
					value: v.value,
				}
				newData[index] = en
			} else {
				for oldEntry.next != nil {
					oldEntry = oldEntry.next
				}
				en := &Entry{
					key:   v.key,
					value: v.value,
				}
				oldEntry.next = en
			}
			v = v.next
		}
	}
	hm.data = newData
}

func getIndex(hash, dataLength int) int {
	return hash % dataLength
}

func hashCode(key interface{}) int {
	bs := make([]byte, 0)
	switch key.(type) {
	case string:
		bs = []byte(key.(string))
		break
	default:
		panic("key type is error")
	}
	c := int(crc32.ChecksumIEEE(bs))
	if c > 0 {
		return c
	}
	if -c > 0 {
		return -c
	}
	return 0
}
func (hm *HashMap) Size() int {
	return hm.size
}
