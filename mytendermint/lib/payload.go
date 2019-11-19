package lib

import (
	"github.com/tendermint/tendermint/crypto"
)

type Payload interface {
	GetType()string
	GetSigner()crypto.Address
	GetSignBytes()[]byte
}

type IssuePayload struct {
	Issuer crypto.Address
	To     crypto.Address
	Value	int64
}
func NewIssuePayload(issuer,to crypto.Address,value int64)*IssuePayload{
	return &IssuePayload{Issuer:issuer,To:to,Value:value}
}

func (pld *IssuePayload) GetType() string{
	return "issueTx"
}

func (pld *IssuePayload) GetSigner() crypto.Address{
	return pld.Issuer
}

func (pld *IssuePayload) GetSignBytes()[]byte{
	bz,err := codec.MarshalJSON(pld)
	if err != nil {
		return []byte{}
	}
	return bz
}

type TxPayload struct{
	From	crypto.Address
	To 		crypto.Address
	Value 	int64
}

func NewTxPayload(from,to crypto.Address,value int64)*TxPayload{
	return &TxPayload{From: from,To:to,Value:value,}
}

func (pld *TxPayload) GetType() string{
	return "transfer"
}

func (pld *TxPayload) GetSigner() crypto.Address{
	return pld.From
}

func (pld *TxPayload) GetSignBytes()[]byte{
	bz,err := codec.MarshalJSON(pld)
	if err != nil {
		return []byte{}
	}
	return bz
}



