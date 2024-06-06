package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	// 区块高度
	Height int64
	//上一个区块哈希
	PrevBlockHash []byte
	//区块数据
	Data []byte
	//时间戳
	Timestamp int64
	//当前区块哈希值
	Hash []byte
	// 难度值
	Nonce int64
}

// 将区块序列化为字节数组
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

// 反序列化
func DeserializeBlock(blockByte []byte) *Block {
	var block Block
	// func gob.NewDecoder(r io.Reader) *gob.Decoder
	decoder := gob.NewDecoder(bytes.NewReader(blockByte))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

// 创建新的区块
func NewBlock(height int64, prevBlockHash []byte, data string) *Block {
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil, 0}
	// 调用工作量证明，返回有效的nonce和hash
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func CreateGenesisBlock(data string) *Block {
	//创世区块
	// func BLC.NewBlock(hight int64, prevBlockHash []byte, data []byte) *BLC.Block
	return NewBlock(1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, data)
}
