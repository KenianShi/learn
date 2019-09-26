package chapter3

import (
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type Blockchain struct {
	tip 	[]byte
	db 		*bolt.DB
}

type BlockchainIterator struct {
	currentHash 	[]byte
	db 				*bolt.DB
}

func (bc *Blockchain) AddBlock(data string){
	var lashHash []byte
	err := bc.db.View(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(blocksBucket))
		lashHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	newBlock := NewBlock(data,lashHash)
	err := bc.db.Update(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash,newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("l"),newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.tip = newBlock.Hash

	})

}