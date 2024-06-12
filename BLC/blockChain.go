package BLC

import (
	"fmt"
	"log"
	"math/big"

	"github.com/boltdb/bolt"
)

type BlockChain struct {
	Tip []byte //最新的区块的Hash
	DB  *bolt.DB
}

// 添加新的区块到链上
// func (blc *BlockChain) AddBlockToBlockChain(data string, height int64, preHash []byte) {
// 	newBlock := NewBlock(height, preHash, data)
// 	// func append(slice []Type, elems ...Type) []Type
// 	blc.Blocks = append(blc.Blocks, newBlock)
// }

// 数据库名字
const dbName = "blockchain.db"

// 表的名字
const blockTableName = "blocks"

// 遍历区块信息：
// 1.验证区块链的完整性和一致性：
// 通过遍历区块链，可以检查每个区块的前一个区块的哈希是否匹配，以确保区块链没有被篡改。

// 2.调试和监控：
// 在开发和测试区块链应用程序时，遍历区块链可以帮助你了解区块链的状态、调试问题和监控链上的数据。

// 3.展示和分析数据：
// 遍历区块链可以用来展示链上的交易数据和其他信息，帮助用户理解区块链的内容。

// 4.数据备份和恢复：
// 在一些情况下，遍历区块链可以用于数据备份和恢复操作，确保链上的数据可以被完整地恢复。
func (blc *BlockChain) PrintChain() {
	var block *Block
	var currentHash []byte = blc.Tip
	for {
		err := blc.DB.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blockTableName))
			if b != nil {
				blockBytes := b.Get(currentHash)
				block = DeserializeBlock(blockBytes)
				fmt.Printf("Height:%d\n", block.Height)
				fmt.Printf("PrevBlockHash:%x\n", block.PrevBlockHash)
				fmt.Printf("Data:%s\n", block.Data)
				fmt.Printf("Timestamp:%d\n", block.Timestamp)
				fmt.Printf("Hash:%x\n", block.Hash)
				fmt.Printf("Nonce:%d\n", block.Nonce)
			}
			return nil
		})
		fmt.Println()
		if err != nil {
			log.Panic(err)
		}
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
		currentHash = block.PrevBlockHash
	}
}

// 增加区块到区块链里
func (blc *BlockChain) AddBlockToBlockchain(data string) {
	// Update(fn func(*bolt.Tx) error) error
	err := blc.DB.Update(func(tx *bolt.Tx) error {
		// 获取区块链数据
		b := tx.Bucket([]byte(blockTableName))
		//创建新区块
		if b != nil {
			//获取最新区块
			blockBytes := b.Get(blc.Tip)
			block := DeserializeBlock(blockBytes)
			//将区块序列化并存储到数据中
			newBlock := NewBlock(block.Height+1, block.Hash, data)
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = b.Put([]byte("1"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			blc.Tip = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

// 创世区块创建
func CreatBlockWithGenesisBlock() *BlockChain {

	// 创建或者打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	var blockHash []byte
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			log.Panic(err)
		}
		if b == nil {
			//创建创世区块
			genesisBlock := CreateGenesisBlock("Genesis Data.......")
			// 将创世区块存储到表中
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			// 存储最新的区块的hash
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			blockHash = genesisBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	// 返回区块链对象
	return &BlockChain{blockHash, db}
}
