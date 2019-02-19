package basic

import (
	"testing"
)

/**
执行顺序，先声明的后执行
*/
func TestDeferStudy(t *testing.T) {

	DeferStudy()
}

func TestRecoverStudy(t *testing.T) {
	t.Log(RecoverStudy())
}

func TestRecoverA(t *testing.T) {
	l, err := RecoverA()
	t.Log("返回结果", l, err)
}

func TestA(t *testing.T) {
	t.Log(A())
	//t.Log(B())
}
