package blockchain

import (
	"fmt"
	"testing"
)

func Test_NewBlockChain(t *testing.T) {

	blockChain := NewBlockChain()
	blockChain.AddBlock("Send 1 BTC TO L")
	blockChain.AddBlock("Send 2 BTC TO R")

	for _, block := range blockChain.blocks {
		fmt.Println("PreHash: ", ConvertToString(string(block.PrevBlockHash), "ISO8859-1", "utf-8"))
		fmt.Println("Data: ", string(block.Data))
		fmt.Println("Hash: ", block.Hash)
	}
}
