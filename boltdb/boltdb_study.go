package boltdb

import (
	"github.com/boltdb/bolt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//以读写的方式打开一个数据库
func OpenBolt(){
	fileDB,err :=  bolt.Open("my.db",os.FileMode(0600),nil)
	if err != nil{
		log.Fatal("open boltdb error:",err)
	}
	defer fileDB.Close()
}
//以读写的方式打开一个数据库，带有超时时间
func OpenBoltOption()  {
	fileDB,err := bolt.Open("my.db",os.FileMode(0600),&bolt.Options{Timeout:1 * time.Second})

	if err != nil{
		log.Fatal("open boltdb error:",err)
	}
	defer fileDB.Close()
}

func BoltUpdate(){
	fileDB,err := bolt.Open("my.db",os.FileMode(0600),&bolt.Options{Timeout:1 * time.Second})
	defer fileDB.Close()
	if err != nil{
		log.Fatal("open boltdb error:",err)
	}
	start := time.Now().Unix()
	fileDB.Update(func(tx *bolt.Tx) error {
		//创建一个bucket,如果存在什么也不做
		bucket,err := tx.CreateBucketIfNotExists([]byte("test.bucket"))
		if err != nil{
			log.Fatal("create bucket error:",err)
		}
		//设置一个随机数种子
		rand.Seed(time.Now().Unix())
		for i:=0;i<1000000 ;i++  {
			bucket.Put([]byte("EMG"+strconv.Itoa(rand.Intn(1000000)+1000000)),[]byte("ss"))
		}
		return nil
	})
	end := time.Now().Unix()

	log.Println(end-start)
}

func BoltBatch(){
	fileDB,err := bolt.Open("my.db",os.FileMode(0600),&bolt.Options{Timeout:1 * time.Second})
	defer fileDB.Close()
	if err != nil{
		log.Fatal("open boltdb error:",err)
	}
	fileDB.Batch(func(tx *bolt.Tx) error {
		//创建一个bucket,如果存在什么也不做
		bucket,err := tx.CreateBucketIfNotExists([]byte("test1.bucket"))
		if err != nil{
			log.Fatal("create bucket error:",err)
		}
		for i:=0;i<1000000 ;i++  {
			bucket.Put([]byte(strconv.Itoa(i)),[]byte("ss"))
		}
		return nil
	})
}