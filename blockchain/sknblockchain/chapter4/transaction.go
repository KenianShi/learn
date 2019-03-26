package chapter4

import (
	"fmt"
	"bytes"
	"encoding/gob"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

const subsidy = 10

type TXOutput struct {
	Value 			int
	ScriptPubKey 	string
}

type TXInput struct {
	Txid 		[]byte
	Vout 		int
	ScrptSig 	string
}

type Transaction struct {
	ID		[]byte
	Vin 	[]TXInput
	Vout 	[]TXOutput
}

func NewUTXOTransaction(from,to string,amount int,bc *Blockchain) *Transaction{
	var inputs []TXInput
	var outputs []TXOutput

	acc,validOutputs := bc.FindSpendableOutputs(from,amount)
	if acc < amount {
		log.Panic("Error:not enough funds")
	}
	for txid,outs := range validOutputs {
		txID,err := hex.DecodeString(txid)
		checkErr(err)
		for _,out := range outs{
			input := TXInput{txID,out,from}
			inputs=append(inputs,input)
		}
	}

	outputs = append(outputs,TXOutput{amount,to})
	if acc > amount{
		outputs = append(outputs,TXOutput{acc - amount,from})
	}
	tx := Transaction{nil,inputs,outputs}
	tx.SetID()
	return &tx
}


func NewCoinbaseTx(to,data string) *Transaction{
	if data == ""{
		data = fmt.Sprintf("Reward to `%s`",to)
	}
	txin := TXInput{[]byte{},-1,data}
	txout := TXOutput{subsidy,to}
	tx := Transaction{nil,[]TXInput{txin},[]TXOutput{txout}}
	tx.SetID()
	return &tx
}

func (bc *Blockchain)FindSpendableOutputs(address string,amount int) (int,map[string][]int){
	unspentOutputs := make(map[string][]int)
	unspentTXs := bc.FindUnspentTransactions(address)
	accmulated := 0
work:
	for _,tx := range unspentTXs{
		txID := hex.EncodeToString(tx.ID)
		for outIndex,out := range tx.Vout{
			if out.CanBeUnlockeedWith(address) && accmulated < amount{
				accmulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID],outIndex)
				if accmulated >= amount{
					break work
				}
			}
		}
	}
	return accmulated,unspentOutputs
}

func (tx *Transaction) SetID(){
	var encoded bytes.Buffer
	var hash [32]byte
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	checkErr(err)
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool{
	return in.ScrptSig == unlockingData
}

func (out *TXOutput) CanBeUnlockeedWith(unlockingData string) bool{
	return out.ScriptPubKey == unlockingData
}

func (tx Transaction) IsCoinbase() bool{
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1			//只有一个输入且输入的txid==0,输入的vout==-1
}


