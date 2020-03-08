package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

/**
声明区块
*/
type Block struct {
	createTime int64
	blockHash  string
	perHash    string
	data       string

	//区块的随机数
	Nonce int
}

/**
创建区块1.0
*/
//func createBlock(perBlock *Block, data string) *Block{
//	newBlock := Block{}
//	newBlock.createTime = time.Now().Unix()
//	newBlock.perHash = perBlock.blockHash
//	newBlock.data = data
//	newBlock.blockHash = operation(&newBlock)
//	return &newBlock
//}

/**
创建区块2.0
增加了工作量证明
*/
func createBlock(perBlockHash string, data string) *Block {
	newBlock := Block{}
	newBlock.createTime = time.Now().Unix()
	newBlock.perHash = perBlockHash
	newBlock.data = data

	pow := NewProofOfWork(&newBlock)
	resultInt, resultByte := pow.Run()
	newBlock.Nonce = resultInt
	newBlock.blockHash = hex.EncodeToString(resultByte)
	return &newBlock
}

/**
区块结构序列化方法
*/
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(result.Bytes())

	var b2 Block
	reder := bytes.NewReader(result.Bytes())
	dereder := gob.NewDecoder(reder)
	dereder.Decode(&b2)

	fmt.Println(b2)
	return result.Bytes()
}

/**
区块链反序列化
*/
func DeserializeBlock(blockBytes []byte) *Block {
	//1.定义一个Block指针对象
	var block Block
	//2.初始化反序列化对象decoder
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	//3.通过Decode()进行反序列化
	err := decoder.Decode(&block)

	if err != nil {
		log.Panic(err)
	}
	//4.返回block对象
	return &block
}

/**
生成hash值
使用工作量证明会弃用此方法
*/
func operation(block *Block) string {
	blockDataStr := string(block.createTime) + block.perHash + block.data
	hashCode := sha256.Sum256([]byte(blockDataStr))
	return hex.EncodeToString(hashCode[:])
}

/**
声明创世区块
*/
func newOneBlock() *Block {
	return createBlock("", "This is foundation Block")
}
