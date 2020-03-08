package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const (
	dbFile      string = "blockchain.db"
	blockBucket string = "blocks"
)

/**
声明区块链结构
*/
//type blockChain struct {
//	blocks []*Block
//}

/**
声明区块链结构2.0
*/
type blockChain struct {
	tip []byte
	Db  *bolt.DB
}

/**
疑问，不知道是什么东西
*/
type BlockchainIterator struct {
	currentHash []byte
	Db          *bolt.DB
}

/**
初始化区块链
*/
func CreationChain() *blockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		//bucket 桶
		b := tx.Bucket([]byte(blockBucket))

		if b == nil {
			fmt.Println("No exising blockchain found Creating a new one ...")
			oneBlock := newOneBlock()

			b, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte(oneBlock.blockHash), oneBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), []byte(oneBlock.blockHash))
			if err != nil {
				log.Panic(err)
			}
			tip = []byte(oneBlock.blockHash)
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := blockChain{tip, db}

	//blockChain := blockChain{}
	//blockChain.blocks = append(blockChain.blocks, newOneBlock())
	return &bc
}

/**
追加区块2.0
改进追加区块方式
*/
func (chainObj *blockChain) AppendBlock(data string) {
	var lastHash []byte

	err := chainObj.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := createBlock(string(lastHash), data)
	fmt.Println("+")
	fmt.Println(newBlock)

	err = chainObj.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		blockBytes := newBlock.Serialize()
		//blockObj := DeserializeBlock(blockBytes)
		//fmt.Println(blockObj)
		err := b.Put([]byte(newBlock.blockHash), blockBytes)

		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), []byte(newBlock.blockHash))
		if err != nil {
			log.Panic(err)
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}

func (chainObj *blockChain) Iterator() {
	chainObj.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		b.ForEach(func(k, v []byte) error {
			if "l" != string(k) {
				block := DeserializeBlock(v)
				fmt.Println(block.perHash)
				fmt.Println(block.createTime)
				fmt.Println(block.data)
				//fmt.Println(string(k))
			}

			return nil
		})
		return nil
	})
}

/**
追加区块1.0
*/
//func (chainObj *blockChain) AppendBlock(data string){
//	newBlock := createBlock(chainObj.blocks[len(chainObj.blocks) - 1], data)
//	chainObj.blocks = append(chainObj.blocks, newBlock)
//}

/**
打印区块链
*/
//func (chainObj *blockChain) PrintChain(){
//	for _, Block := range chainObj.blocks {
//		fmt.Println("Block perHash = ", Block.perHash)
//		fmt.Println("Block createTime = ", Block.createTime)
//		fmt.Println("Block data = ", Block.data)
//		fmt.Println("Block blockHash = ", Block.blockHash)
//
//		//对区块进行工作量证明
//		pow := NewProofOfWork(Block)
//		fmt.Printf("Block is isVerify %v \n", strconv.FormatBool(pow.isVerify()))
//	}
//}
