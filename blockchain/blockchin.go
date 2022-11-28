package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

func (b *block) getHash() {
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

func crateblock(data string) *block {
	newblock := block{data, "", getLasthash()}
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

func AllBlocks() []*block {
	return GetBlockChain().blocks
}
