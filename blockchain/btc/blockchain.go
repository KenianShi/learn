package btc

import (
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const lastest="latest"

type Blockchain struct {
	Tip 	[]byte
	Db 		*bolt.DB
}

func checkErr(msg string,err error){
	if err != nil {
		log.Fatal(msg +"失败，",err)
	}
}

func  NewBlockchain() *Blockchain{
	var tip []byte
	db,err := bolt.Open(dbFile,0600,nil)
	checkErr("打开数据库",err)
	err = db.Update(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			genesis := NewGenesisBlock()
			b,err := tx.CreateBucket([]byte(blocksBucket))
			checkErr("创建blockBucket",err)
			err = b.Put([]byte(genesis.Hash),genesis.Serialize())
			checkErr("添加创世块",err)
			err = b.Put([]byte(lastest),[]byte(genesis.Hash))
			checkErr("",err)
			tip = genesis.Hash
		}else {
			tip = b.Get([]byte(lastest))
		}
		return nil
	})
	bc := Blockchain{Tip:tip,Db:db}
	return &bc
}

func (bc *Blockchain) AddBlock(data string){
	var latestHash []byte
	err := bc.Db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		latestHash = bucket.Get([]byte(lastest))
		return nil
	})
	checkErr("查看最新去块，",err)
	newBlock := NewBlock(data,latestHash)
	err = bc.Db.Update(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put([]byte(newBlock.Hash),newBlock.Serialize())
		checkErr("添加区块",err)
		err = b.Put([]byte(lastest),newBlock.Hash)
		checkErr("更新区块，",err)
		bc.Tip = newBlock.Hash
		return nil
	})
}

type BlockchainIterator struct {
	currentHash 	[]byte
	db 				*bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator{
	bci := &BlockchainIterator{bc.Tip,bc.Db}
	return bci
}

func (i *BlockchainIterator) Next() *Block{
	var block *Block
	err := i.db.View(func(tx *bolt.Tx)error{
		bucket := tx.Bucket([]byte(blocksBucket))
		encodeBlock := bucket.Get([]byte(i.currentHash))
		block = DeserializeBlock(encodeBlock)
		return nil
	})
	checkErr("查询区块",err)
	i.currentHash = block.PreHash
	return block
}
