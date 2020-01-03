package basic

import "testing"

//sync 包的学习

func TestSyncMap(t *testing.T) {
	for i := 0; i < 100; i++ {
		SyncMapPut(i, i)
	}
	for i := 150; i < 200; i++ {
		SyncMapUpdate(i, i)
	}
	for i := 0; i < 100; i++ {
		SyncMapRead(i)
	}

	SyncMapRange()
}

func TestStudyMutex(t *testing.T) {
	StudyMutex()
}

func TestStudySyncOnce(t *testing.T) {
	StudySyncOnce()
}

func TestStudySyncPool(t *testing.T) {
	StudySyncPool()

}
