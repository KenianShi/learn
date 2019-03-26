package chapter4

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"os"
)

const dbFile = "blockchain4.db"
const blocksBucket = "block"
const genesisCoinbaseData = "shikenian's first block"
const latestHashTag = "latest"

type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func CreateBlockchain(address string) *Blockchain {
	if dbExits() {
		fmt.Println("blockchain already exits")
		os.Exit(1)
	}
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	checkErr(err)
	err = db.Update(func(tx *bolt.Tx) error {
		cbtx := NewCoinbaseTx(address, genesisCoinbaseData)
		genesis := NewGenesisBlock(cbtx)
		b, err := tx.CreateBucket([]byte(blocksBucket))
		checkErr(err)
		err = b.Put(genesis.Hash, genesis.Serialize())
		checkErr(err)
		err = b.Put([]byte(latestHashTag), genesis.Hash)
		checkErr(err)
		tip = genesis.Hash
		return nil
	})
	checkErr(err)
	bc := Blockchain{tip, db}
	return &bc
}

func dbExits() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction {
	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	for {
		block := bci.Next()
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		outputs:
			for outIndex, out := range tx.Vout {
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if outIndex == spentOut {
							continue outputs
						}
					}
				}
				if out.CanBeUnlockeedWith(address) {
					unspentTXs = append(unspentTXs, *tx) //这儿可能有一个问题，放这笔交易的输出被后面的交易引用的时候，这里并没有剔除掉
				}
			}
			if !tx.IsCoinbase() {
				for _, in := range tx.Vin {
					if in.CanUnlockOutputWith(address) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout) //同上的注释

					}
				}
			}
		}
		if len(block.PreBlockHash) == 0 {
			break
		}
	}
	return unspentTXs
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}
	return bci
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block
	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)
		return nil
	})
	checkErr(err)
	i.currentHash = block.PreBlockHash
	return block
}

func (bc *Blockchain) FindUTXO(address string) []TXOutput{
	var UTXOs []TXOutput
	unspentTransactions := bc.FindUnspentTransactions(address)
	for _,tx := range unspentTransactions{
		for _,out := range tx.Vout{
			if out.CanBeUnlockeedWith(address){
				UTXOs = append(UTXOs,out)
			}
		}
	}
	return UTXOs
}

func (bc *Blockchain) MineBlock(transactions []*Transaction){
	var lastHash []byte
	err := bc.db.View(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte(latestHashTag))
		return nil
	})
	checkErr(err)

	newBlock := NewBlock(transactions,lastHash)
	err = bc.db.Update(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(blocksBucket))
		err = b.Put(newBlock.Hash,newBlock.Serialize())
		checkErr(err)
		err = b.Put([]byte(latestHashTag),newBlock.Hash)
		checkErr(err)
		bc.tip = newBlock.Hash
		return nil
	})
	checkErr(err)
}
