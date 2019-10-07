package register_center

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc/resolver"
)

type etcdBuilder struct {
	address     string
	client      *clientv3.Client
	serviceName string
}

func NewEtcdBuilder(address string) resolver.Builder {

	client, err := clientv3.NewFromURL(address)
	if err != nil {
		log.Fatal("LearnGrpc: create etcd client error", err.Error())
		return nil
	}
	return &etcdBuilder{address: address, client: client}
}

func (cb *etcdBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	cb.serviceName = target.Endpoint

	adds, serviceConfig, err := cb.resolve()
	if err != nil {
		return nil, err
	}
	cc.NewAddress(adds)
	cc.NewServiceConfig(serviceConfig)
	etcdResolver := NewEtcdResolver(&cc, cb, opts)
	etcdResolver.wg.Add(1)
	go etcdResolver.watcher()

	return etcdResolver, nil
}

func (cb etcdBuilder) resolve() ([]resolver.Address, string, error) {

	res, err := cb.client.Get(context.Background(), cb.serviceName)
	if err != nil {
		return nil, "", err
	}
	adds := make([]resolver.Address, 0)
	for i := range res.Kvs {
		if v := res.Kvs[i].Value; v != nil {
			temp := resolver.Address{Addr: string(v), ServerName: cb.serviceName}
			adds = append(adds, temp)
		}
	}
	return adds, "", nil
}

func (cb *etcdBuilder) Scheme() string {
	return "etcd"
}

type etcdResolver struct {
	clientConn           *resolver.ClientConn
	etcdBuilder          *etcdBuilder
	t                    *time.Ticker
	wg                   sync.WaitGroup
	rn                   chan struct{}
	ctx                  context.Context
	cancel               context.CancelFunc
	disableServiceConfig bool
}

func NewEtcdResolver(cc *resolver.ClientConn, cb *etcdBuilder, opts resolver.BuildOption) *etcdResolver {
	ctx, cancel := context.WithCancel(context.Background())
	return &etcdResolver{
		clientConn:           cc,
		etcdBuilder:          cb,
		t:                    time.NewTicker(time.Second),
		ctx:                  ctx,
		cancel:               cancel,
		disableServiceConfig: opts.DisableServiceConfig}
}

func (cr *etcdResolver) watcher() {
	cr.wg.Done()
	for {
		select {
		case <-cr.ctx.Done():
			return
		case <-cr.rn:
		case <-cr.t.C:
		}
		adds, serviceConfig, err := cr.etcdBuilder.resolve()
		if err != nil {
			log.Fatal("query service entries error:", err.Error())
		}
		(*cr.clientConn).NewAddress(adds)
		(*cr.clientConn).NewServiceConfig(serviceConfig)
	}
}

func (cr *etcdResolver) Scheme() string {
	return cr.etcdBuilder.Scheme()
}

func (cr *etcdResolver) ResolveNow(rno resolver.ResolveNowOption) {
	select {
	case cr.rn <- struct{}{}:
	default:
	}
}

func (cr *etcdResolver) Close() {
	cr.cancel()
	cr.wg.Wait()
	cr.t.Stop()
}

type etcdClientConn struct {
	adds  []resolver.Address
	sc    string
	state resolver.State
}

func NewEtcdClientConn() resolver.ClientConn {
	return &etcdClientConn{}
}

func (cc *etcdClientConn) NewAddress(addresses []resolver.Address) {
	cc.adds = addresses
}

func (cc *etcdClientConn) NewServiceConfig(serviceConfig string) {
	cc.sc = serviceConfig
}

func (cc *etcdClientConn) UpdateState(state resolver.State) {
	cc.state = state
}

func GenerateAndRegisterEtcdResolver(address string, serviceName string) (schema string, err error) {
	builder := NewEtcdBuilder(address)
	target := resolver.Target{Scheme: builder.Scheme(), Endpoint: serviceName}
	_, err = builder.Build(target, NewEtcdClientConn(), resolver.BuildOption{})
	if err != nil {
		return builder.Scheme(), err
	}
	resolver.Register(builder)
	schema = builder.Scheme()
	return schema, nil
}
