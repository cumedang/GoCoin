package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func (b *Block) getHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLasthash() string {
	totalBlocks := len(GetBlockChain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockChain().blocks[totalBlocks-1].Hash
}

func crateblock(data string) *Block {
	newblock := Block{data, "", getLasthash()}
	newblock.getHash()
	return &newblock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, crateblock(data))
}

func GetBlockChain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Gensis")
		})
	}
	return b
}

func AllBlocks() []*Block {
	return GetBlockChain().blocks
}
