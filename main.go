package main

import (
	"fmt"
	blockchain "github.com/genkovich/go-test-blockchain"
)

func main() {
	chain := blockchain.CreateBlockchain(3)

	chain.AddBlock("Test", "Target", 5)
	chain.AddBlock("John", "Bob", 2)

	fmt.Println(chain.IsValid())
	fmt.Println(chain)
}
