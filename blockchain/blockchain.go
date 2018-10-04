package blockchain

import (
	"bytes"
	"crypto/sha256"
	"github.com/axgle/mahonia"
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

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{},
	}
	block.SetHash()
	return block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func (blockChain *BlockChain) AddBlock(data string) {
	preBlock := blockChain.blocks[len(blockChain.blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	blockChain.blocks = append(blockChain.blocks, newBlock)
}

/**
创始快
*/
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block创始块", []byte{})
}

/**
用创始块创建一个区块链函数
*/
func NewBlockChain() *BlockChain {
	return &BlockChain{
		[]*Block{NewGenesisBlock()},
	}
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
