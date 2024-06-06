package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 256位哈希里面前面至少需要有16个0
const targetBit = 16

type ProofOfWork struct {
	// 当前需要验证的区块
	Block *Block
	// 生成有效哈希所需的难度目标
	target *big.Int
}

// 数据拼接，输入nonce
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.Block.PrevBlockHash,
		pow.Block.Data,
		IntToHex(pow.Block.Timestamp),
		IntToHex(int64(targetBit)),
		IntToHex(int64(nonce)),
		IntToHex(int64(pow.Block.Height)),
	}, []byte{})
	return data
}

func (proofOfWork *ProofOfWork) IsValid() bool {
	// 存储哈希值
	var hashInt big.Int
	hashInt.SetBytes(proofOfWork.Block.Hash)
	// Cmp(y *big.Int) (r int)
	return proofOfWork.target.Cmp(&hashInt) == 1
}

// 启动工作量证明
func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	nonce := 0
	// 存储哈希值
	var hashInt big.Int
	// 存储256字节长度的哈希值
	var hash [32]byte
	for {
		dataBytes := proofOfWork.prepareData(nonce)
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		// func (z *big.Int) SetBytes(buf []byte) *big.Int
		hashInt.SetBytes(hash[:])
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce = nonce + 1
	}
	return hash[:], int64(nonce)
}

// 创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	// 难度初始值为1
	target := big.NewInt(1)
	// 难度值小，挖矿难度越难
	target = target.Lsh(target, 256-targetBit)
	return &ProofOfWork{block, target}
}
