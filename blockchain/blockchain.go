package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

/**
一个简单的区块链
*/
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

type BlockChain struct {
	blocks []*Block
}

func newBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{},
	}
	block.setHash()
	return block
}

/**
给区块设置hash值
*/
func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

/**
添加一个新的区块
*/
func (blockChain *BlockChain) AddBlock(data string) {
	preBlock := blockChain.blocks[len(blockChain.blocks)-1]
	newBlock := newBlock(data, preBlock.Hash)
	blockChain.blocks = append(blockChain.blocks, newBlock)
}

/**
创始快
*/
func NewGenesisBlock() *Block {
	return newBlock("Genesis Block创始块", []byte{})
}

/**
用创始块创建一个区块链函数
*/
func NewBlockChain() *BlockChain {
	return &BlockChain{
		[]*Block{NewGenesisBlock()},
	}
}
