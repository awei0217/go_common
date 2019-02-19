package basic

import (
	"fmt"
	"sync"
)

var m sync.Map

/**
添加数据
 */
func SyncMapPut(key int, value interface{}){
	m.Store(key,value)
}

/**
返回键的现有值(如果存在)，否则存储并返回给定的值，如果是读取则返回true，如果是存储返回false。
 */
func SyncMapUpdate(key int,value interface{}){
	actual,ok := m.LoadOrStore(key,value)
	fmt.Println(actual,ok)

}
/**
删除指定的key
 */
func SyncMapDel(key int){
	m.Delete(key)
}

/**
读取key的值，如果不存在，返回 nil,false
 */
func SyncMapRead(key int){
	v,ok := m.Load(key)
	fmt.Println(v,ok)
}

func SyncMapRange(){
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}


