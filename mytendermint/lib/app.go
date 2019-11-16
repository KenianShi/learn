package lib

import (
	"bytes"
	"errors"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"math/big"
)

var (
	SYSTEM_ISSUER = crypto.Address("KING_OF_TOKEN")
)

type TokenApp struct {
	types.BaseApplication
	Accounts 	map[string]*big.Int
}

func NewTokenApp() *TokenApp{
	return &TokenApp{Accounts: map[string]*big.Int{}}
}

func (app *TokenApp) CheckTx(req types.RequestCheckTx)(rsp types.ResponseCheckTx){
	tx,err := app.decodeTx(req.Tx)
	if err != nil {
		rsp.Code = 1
		rsp.Log = "decode tx error"
		return
	}
	if !tx.Verify(){
		rsp.Code = 2
		rsp.Log = "tx verify failed"
		return
	}
	return
}

func (app *TokenApp) DeliverTx(req types.RequestDeliverTx)(rsp types.ResponseDeliverTx){
	tx,_ := app.decodeTx(req.Tx)				//如果err，在之前的checkTx就通过不了，故此处可以省略
	switch tx.Payload.GetType(){
	case "issueTx":
		pld := tx.Payload.(*IssuePayload)
		err := app.issue(pld.Issuer,pld.To,pld.Value)
		if err != nil {
			rsp.Log = err.Error()
		}
		rsp.Info = "issuer tx applied!"
	case "transfer":
		pld := tx.Payload.(*TxPayload)
		err := app.transfer(pld.From,pld.To,pld.Value)
		if err != nil {
			rsp.Log = err.Error()
		}
		rsp.Info = "transfer tx applied"
	}
	return
}

func (app *TokenApp) Query(req types.RequestQuery)(rsp types.ResponseQuery){
	add := crypto.Address(req.Data)
	rsp.Key = req.Data
	rsp.Value,_ = codec.MarshalJSON(app.Accounts[add.String()])
	return
}



func (app *TokenApp) decodeTx(bz []byte)(*Tx,error){
	var tx Tx
	err := codec.UnmarshalJSON(bz,&tx)
	return &tx,err
}

func (app *TokenApp) issue(issuer,to crypto.Address,value *big.Int)error{
	wallet := LoadWallet()
	SYSTEM_ISSUER = wallet.GetAddress("issuer")
	if !bytes.Equal(issuer,SYSTEM_ISSUER){
		return errors.New("invalid issuer")
	}
	app.Accounts[to.String()].Add(app.Accounts[to.String()],value)
	return nil
}

func (app *TokenApp) transfer(from,to crypto.Address,value *big.Int)error{
	if app.Accounts[from.String()].Cmp(value) < 0 {
		return errors.New("not enough balance")
	}
	app.Accounts[from.String()].Sub(app.Accounts[from.String()],value)
	app.Accounts[to.String()].Add(app.Accounts[to.String()],value)
	return nil
}

