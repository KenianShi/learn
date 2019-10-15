package lab

import (
	"bytes"
	"github.com/tendermint/tendermint/crypto"
	"time"
)

type Tx struct {
	Payload		Payload
	Signature 	[]byte
	Pubkey 		crypto.PubKey
	Sequence 	int64
}

func NewTx(payload Payload) *Tx{
	return &Tx{
		Payload:   payload,
		Sequence:  time.Now().Unix(),
	}
}

func (tx *Tx) Sign(priv crypto.PrivKey)error{
	data := tx.Payload.GetSignBytes()
	var err error
	tx.Signature,err = priv.Sign(data)				//对整个payload进行签名，签名结果赋值给tx的signature属性值
	tx.Pubkey = priv.PubKey()
	return err
}


func (tx *Tx) Verify() bool{
	signer := tx.Payload.GetSigner()
	signerFromKey := tx.Pubkey.Address()
	if !bytes.Equal(signer,signerFromKey){
		return false
	}
	data := tx.Payload.GetSignBytes()
	sig := tx.Signature
	valid := tx.Pubkey.VerifyBytes(data,sig)
	if !valid{
		return false
	}
	return true
}





