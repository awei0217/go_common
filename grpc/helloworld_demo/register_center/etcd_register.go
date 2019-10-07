package register_center

import (
	"context"
	"fmt"
	"log"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
)

type etcdRegisterImpl struct {
	Address string
}

//创建etcd的注册结构体
func NewEtcdRegisterImpl(target string) *etcdRegisterImpl {
	return &etcdRegisterImpl{Address: target}
}

//etcd 实现注册接口
func (etcd *etcdRegisterImpl) Register(info ServiceDescInfo) error {

	client, err := clientv3.NewFromURL(etcd.Address)
	if err != nil {
		log.Println("连接etcd失败:", err)
		return err
	}
	// minimum lease TTL is ttl-second
	resp, err := client.Grant(context.TODO(), int64(info.IntervalTime))
	if err != nil {
		log.Println("创建租约失败:", err)
		return err
	}
	// should get first, if not exist, set it
	_, err = client.Get(context.Background(), info.ServiceName)
	serviceValue := fmt.Sprintf("%s:%d", info.Host, info.Port)
	if err != nil {
		if err == rpctypes.ErrKeyNotFound {
			if _, err := client.Put(context.TODO(), info.ServiceName, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
				log.Printf("etcd: set service '%s' with ttl to etcd3 failed: %s", info.ServiceName, err.Error())
			}
		} else {
			log.Printf("etcd: service '%s' connect to etcd3 failed: %s", info.ServiceName, err.Error())
			return err
		}
	} else {
		// refresh set to true for not notifying the watcher
		if _, err := client.Put(context.Background(), info.ServiceName, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
			log.Printf("etcd: refresh service '%s' with ttl to etcd3 failed: %s", info.ServiceName, err.Error())
			return err
		}
	}
	log.Println("register successful")
	return nil
}

//etcd 实现下线接口
func (etcd *etcdRegisterImpl) UnRegister(info ServiceDescInfo) error {
	client, err := clientv3.NewFromURL(etcd.Address)
	if _, err := client.Delete(context.Background(), info.ServiceName); err != nil {
		log.Printf("etcd: deregister '%s' failed: %s", info.ServiceName, err.Error())
	} else {
		log.Printf("etcd: deregister '%s' ok.", info.ServiceName)
	}
	return err
}
