package algorithm

import "testing"

func TestNewWeightedRoundRobin(t *testing.T) {
	wrr := NewWeightedRoundRobin(map[string]int{"10.1": 7, "10.2": 2, "10.3": 1})
	for i := 0; i < 10; i++ {
		sw := wrr.Select()
		t.Log(sw.Ip)
	}

}
