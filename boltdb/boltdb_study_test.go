package boltdb

import (
	"testing"
)

func TestOpenBolt(t *testing.T) {
	OpenBolt()
}
//370.10s
func TestBoltUpdate(t *testing.T) {
	BoltUpdate()
}
//356.36s
func TestBoltBatch(t *testing.T) {
	BoltBatch()
}