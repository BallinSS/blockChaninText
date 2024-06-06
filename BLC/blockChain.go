package BLC

import (
	"log"

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

// 创世区块创建
func CreatBlockWithGenesisBlock() *BlockChain {

	// 创建或者打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
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
