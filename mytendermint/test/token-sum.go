package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/tendermint/tendermint/crypto"
)

var (
	SYSTEM_ISSUER1 = crypto.Address("KING_OF_TOKEN")
)

type TokenApp1 struct {
	Accounts map[string]int
}

func NewTokenApp1()*TokenApp1{
	return &TokenApp1{Accounts: map[string]int{}}
}

func (app *TokenApp1)transfer(from,to crypto.Address,value int)error{
	if app.Accounts[from.String()] < value {
		return errors.New("not enough balance")
	}
	app.Accounts[from.String()] -= value
	app.Accounts[to.String()] += value
	return nil
}

func (app *TokenApp1) issue(issuer,to crypto.Address,value int)error{
	if !bytes.Equal(issuer,SYSTEM_ISSUER1){
		return errors.New("invalid issuer")
	}
	app.Accounts[to.String()] += value
	return nil
}

func (app *TokenApp1) Dump1(){
	fmt.Printf("state => %v\n",app.Accounts)
}


func main() {
	app := NewTokenApp1()
	a1 := crypto.Address("TEST_ADDRESS1")
	a2 := crypto.Address("TEST_ADDRESS2")
	app.issue(SYSTEM_ISSUER1,a1,200)
	app.transfer(a1,a2,20)
	app.Dump1()

}

