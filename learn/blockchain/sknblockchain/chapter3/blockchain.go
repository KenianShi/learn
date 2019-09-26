package chapter3

import (
	"github.com/boltdb/bolt"
	"fmt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type Blockchain struct {
	Tip			[]byte
	Db 			*bolt.DB
}

type BlockchainIterator struct {
	currentHash 		[]byte
	db 					*bolt.DB
}

func NewBlockchain() *Blockchain{
	var tip []byte
	db,err := bolt.Open(dbFile,0600,nil)
	checkErr(err)
	err = db.Update(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil{
			fmt.Println("No exiting blockchain found.Creating a new one ...")
			genesis := NewGenesisBlock()
			b,err := tx.CreateBucket([]byte(blocksBucket))
			checkErr(err)
			err = b.Put(genesis.Hash,genesis.Serialize())
			checkErr(err)

			err = b.Put([]byte("latest"),genesis.Hash)
			checkErr(err)

			tip = genesis.Hash

		}else {
			tip = b.Get([]byte("latest"))
		}
		return nil
	})
	checkErr(err)

	bc := &Blockchain{Tip:tip,Db:db}
	return bc
}

func (bc *Blockchain) AddBlock(data string){
	var lastHash []byte
	err := bc.Db.View(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("latest"))
		return nil
	})
	checkErr(err)
	newBlock := NewBlock(data,lastHash)
	err = bc.Db.Update(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash,newBlock.Serialize())
		checkErr(err)
		err = b.Put([]byte("latest"),newBlock.Hash)
		checkErr(err)
		bc.Tip = newBlock.Hash
		return nil
	})
}

func (bc *Blockchain) Iterator() *BlockchainIterator{
	bci := &BlockchainIterator{bc.Tip,bc.Db}
	return bci
}

func (i *BlockchainIterator) Next() *Block{
	var block *Block
	err := i.db.View(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)
		return nil
	})
	checkErr(err)
	i.currentHash = block.PrevBlockHash
	return block
}