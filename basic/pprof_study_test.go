package basic

import (
	"testing"

	_ "net/http/pprof"
)

// 针对web性能监测 ,引入 _ "net/http/pprof" 就可以了

func TestPprofWeb(t *testing.T) {

	PprofWeb()

}

func TestPProfCPUApplication(t *testing.T) {
	PProfCPUApplication()
}

func TestPProfMemApplication(t *testing.T) {
	PProfMemApplication()
}
