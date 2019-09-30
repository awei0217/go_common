package client

import (
	"testing"
)

func TestStartClient(t *testing.T) {
	StartClient()
}

func BenchmarkStartClient(b *testing.B) {
	StartClient()
}
