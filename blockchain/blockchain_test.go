package blockchain

import (
	"fmt"
	"testing"
)

func Test_NewBlockChain(t *testing.T) {

	blockChain := NewBlockChain()
	blockChain.AddBlock("添加第一个区块")
	blockChain.AddBlock("添加第二个区块")

	for _, block := range blockChain.blocks {
		fmt.Println("PreHash: ", block.PrevBlockHash)
		fmt.Println("Data: ", string(block.Data))
		fmt.Println("Hash: ", block.Hash)
	}
}
