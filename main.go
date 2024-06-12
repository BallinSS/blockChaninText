package main

import (
	"blc/BLC"
)

func main() {
	blockChain := BLC.CreatBlockWithGenesisBlock()
	defer blockChain.DB.Close()

	blockChain.AddBlockToBlockchain("send 100BTC to Ballin")
}
