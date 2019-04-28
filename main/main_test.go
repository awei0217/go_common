package main

import (
	"io"
	"testing"
)

func TestQueryHelloWorld(t *testing.T) {

	result, err := QueryHelloWorld("test")

	if err != nil {
		t.Log(result)
	}
	t.Log(err)

}

type Hash interface {
	io.Writer

	Sum(b []byte) []byte

	Reset()

	Size() int

	BlockSize() int
}
