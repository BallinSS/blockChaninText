package main

import (
	"blc/BLC"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	//创建打开数据库
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		//立即停止
		log.Fatal(err)
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		if b != nil {
			blockData := b.Get([]byte("l"))
			block := BLC.DeserializeBlock(blockData)
			fmt.Printf("%v\n", block)
		}
		return nil
	})

	if err != nil {
		// 执行defer后停止
		log.Panic(err)
	}
	// 初始化区块链
	// block := BLC.NewBlock(1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, "Test")
	// fmt.Printf("%d\n", block.Nonce)
	// fmt.Printf("%x\n", block.Hash)
	// bytes := block.Serialize()

	// fmt.Println(bytes)

	// block = BLC.DeserializeBlock(bytes)

	// fmt.Printf("%d\n", block.Nonce)
	// fmt.Printf("%x\n", block.Hash)
}
