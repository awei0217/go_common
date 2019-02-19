package basic

import "testing"

func TestSyncMap(t *testing.T) {
	for i:=0;i<100;i++  {
		SyncMapPut(i,i)
	}
	for i:=150;i<200;i++  {
		SyncMapUpdate(i,i)
	}
	for i:=0;i<100;i++  {
		SyncMapRead(i)
	}

	SyncMapRange()
}



