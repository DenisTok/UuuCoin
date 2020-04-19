package main

import (
	"UuuCoin/src/block"
)

func main() {
	firsBlock := block.NewBlock(&block.Block{}, block.NewData([]byte("Hello, world")))
	secondBlock := block.NewBlock(firsBlock, block.NewData([]byte("Hello, world 2")))
	firsBlock.Print()
	secondBlock.Print()
}
