package core

import (
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 16

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

//构建工作证明结构体变量
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := ProofOfWork{b, target}
	return &pow
}

//随机数转换
func (pow *ProofOfWork) preparData(nonce int) []byte {
	data := pow.block.perHash + pow.block.data + string(pow.block.createTime) + string(targetBits) + string(nonce)
	return []byte(data)
}

func (pow *ProofOfWork) isVerify() bool {
	var hashInt big.Int
	resultBlockStr := pow.preparData(pow.block.Nonce)
	resultHash := sha256.Sum256(resultBlockStr)
	hashInt.SetBytes(resultHash[:])
	//if hex.EncodeToString(resultHash[:]) == pow.Block.blockHash {
	//	return true
	//}
	return hashInt.Cmp(pow.target) == -1
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the Block containing \" %v \"\n ", pow.block.data)
	for nonce < maxNonce {
		data := pow.preparData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println("\n")
	return nonce, hash[:]
}
