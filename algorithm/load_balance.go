package algorithm

type SeverWeight struct {
	//配置的权重
	ConfigWeight int
	//当前权重
	CurrentWeight int
	//有效权重
	EffectiveWeight int
	//服务器ip
	Ip string
}

//加权轮询算法
type WeightedRoundRobin struct {
	//机器ip和对应的权重
	IpAndWeightedConfig map[string]int
	//服务器和权重信息
	SwSlice []*SeverWeight
}

func NewWeightedRoundRobin(iwc map[string]int) *WeightedRoundRobin {
	if iwc == nil {
		return nil
	}
	SwSlice := make([]*SeverWeight, 0)
	for k, v := range iwc {
		sw := &SeverWeight{ConfigWeight: v, CurrentWeight: 0, EffectiveWeight: v, Ip: k}
		SwSlice = append(SwSlice, sw)
	}
	return &WeightedRoundRobin{IpAndWeightedConfig: iwc, SwSlice: SwSlice}
}
func (wrr *WeightedRoundRobin) Select() (sw *SeverWeight) {
	total := 0
	for _, v := range wrr.SwSlice {
		v.CurrentWeight += v.EffectiveWeight
		total += v.EffectiveWeight
		if v.EffectiveWeight < v.ConfigWeight {
			v.EffectiveWeight++
		}
		if sw == nil || v.CurrentWeight > sw.CurrentWeight {
			sw = v
		}
	}
	sw.CurrentWeight = sw.CurrentWeight - total
	return sw
}

//一致性hash算法
type ConsistencyHash struct {
}
