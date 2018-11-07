package redis

import "testing"

func TestString(t *testing.T) {
	String()
}

func TestHash(t *testing.T) {
	Hash()
}

func TestList(t *testing.T) {
	List()
}

func TestSet(t *testing.T) {
	Set()
}

func TestSetDistributeLock(t *testing.T) {
	SetDistributeLock()
}

func TestReleaseDistributeLock(t *testing.T) {
	ReleaseDistributeLock()
}

func TestSortedSet(t *testing.T) {
	SortedSet()
}

func TestDelayedQueue(t *testing.T) {
	DelayedQueue()
}

func TestHyperLogLog(t *testing.T) {
	HyperLogLog()
}

func TestBitMap(t *testing.T) {
	BitMap()
}

func TestBloomFilter(t *testing.T) {
	BloomFilter()
}