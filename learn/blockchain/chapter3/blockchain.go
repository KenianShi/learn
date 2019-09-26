package chapter3

import (
	"github.com/boltdb/bolt"
)

const dbFile = "./db/chapter3db"
type Blockchain struct {
	Tip []byte
	Db  *bolt.DB
}

func (bc *Blockchain) AddBlock(data string){
	var lastHash []byte
	err := bc.Db.View(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(BlockBuckerName))
		lastHash = b.Get([]byte(LatestHashtag))
		return nil
	})
	if err != nil {

	}
	newBlock := NewBlock(data,lastHash)
	err = bc.Db.Update(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(BlockBuckerName))
		err := b.Put(newBlock.Hash,newBlock.Serialize())
		if err != nil {

		}
		err = b.Put([]byte(LatestHashtag),newBlock.Hash)
		if err != nil {

		}
		bc.Tip = newBlock.Hash
		return nil
	})
	if err != nil {

	}
}

func NewBlockchain() *Blockchain{
	var tip []byte
	db,err := bolt.Open(dbFile,0600,nil)
	if err != nil {

	}
	err = db.Update(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(BlockBuckerName))
		if b == nil {
			genesis := NewGenesisBlock()
			b,err := tx.CreateBucket([]byte(BlockBuckerName))
			if err != nil {

			}
			err = b.Put(genesis.Hash,genesis.Serialize())
			if err != nil {

			}
			err = b.Put([]byte(LatestHashtag),genesis.Hash)
			if err != nil {

			}
			tip = genesis.Hash
		}else{
			tip = b.Get([]byte(LatestHashtag))
		}
		return nil
	})
	bc := Blockchain{tip,db}
	return &bc
}

type BlockchainIterator struct {
	currentHash		[]byte
	db 				*bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator{
	bci := &BlockchainIterator{bc.Tip,bc.Db}
	return bci
}

func (i *BlockchainIterator) Next() *Block{
	var block *Block
	err := i.db.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte(BlockBuckerName))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)
		return nil
	})
	if err != nil {

	}
	i.currentHash = block.Prehash
	return block
}
