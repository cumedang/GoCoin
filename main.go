package main

import (
	"fmt"

	"github.com/cumedang/GoCoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")
	for _, block := range blockchain.AllBlocks() {
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %s\n", block.Hash)
		fmt.Printf("Prev Hash : %s\n", block.PrevHash)
	}
}
