package main

import (
	"fmt"
	"strconv"

	"github.com/wadro/golang-study/blockchain/bc"
)

/*
Tools environment: GOPATH=C:\Users\echo1\go
Installing 3 tools at C:\Users\echo1\go\bin in module mode.
  go-outline
  dlv
  staticcheck
*/

func main() {
	chain := bc.InitBlockChain()
	chain.AddBlock("First")
	chain.AddBlock("Second")
	chain.AddBlock("Third")

	for _, block := range chain.Blocks {
		fmt.Printf("\nPrev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)

		pow := bc.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
	fmt.Println(chain)
}
