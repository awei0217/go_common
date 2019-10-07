package register_center

import "time"

//服务描述信息
type ServiceDescInfo struct {
	//服务名称
	ServiceName string
	//ip地址
	Host string
	//端口
	Port int
	//心跳间隔 秒
	IntervalTime time.Duration
}

//服务注册和下线的接口
type RegisterI interface {

	//服务注册
	Register(serviceInfo ServiceDescInfo) error

	//服务下线
	UnRegister(serviceInfo ServiceDescInfo) error
}
