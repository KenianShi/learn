package lib

import (
	"bytes"
	"fmt"
	"github.com/tendermint/tendermint/crypto"
	"time"
)

type Tx struct {
	Payload 	Payload				`json:"payload"`
	PubKey 		crypto.PubKey		`json:"pub_key"`
	Signature 	[]byte				`json:"signature"`
	Sequence 	int64				`json:"sequence"`
}

type Transaction interface {
	Sign(crypto.PrivKey) error
	Verify() bool
}


func NewTransaction(pld Payload) Tx{
	sequence := time.Now().Unix()
	return Tx{Payload:pld,Sequence:sequence}
}

func (tx *Tx) Sign(privKey crypto.PrivKey) error{
	data := tx.Payload.GetSignBytes()
	signature,err := privKey.Sign(data)
	if err != nil{
		panic(err)
	}
	tx.PubKey = privKey.PubKey()
	tx.Signature = signature
	return err
}

func (tx *Tx) Verify() bool {
	signer := tx.Payload.GetSigner()
	addressFromKey := tx.PubKey.Address()
	if !bytes.Equal(signer,addressFromKey) {
		fmt.Println("Signer and the addressFromKey not matched")
		return false
	}
	data := tx.Payload.GetSignBytes()
	signature := tx.Signature
	//fmt.Printf("打印信息3：\n%v\n%v\n",data,signature)
	return tx.PubKey.VerifyBytes(data,signature)
}

