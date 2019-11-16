package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/KenianShi/learn/mytendermint/lib"
	"github.com/tendermint/tendermint/crypto"
	kf "github.com/tendermint/tendermint/crypto/secp256k1"
	"math/big"
)

var (
	issuer		= kf.GenPrivKey()
	SYSTEM_ISSUER = issuer.PubKey().Address()
)

type TokenApp struct {
	Accounts map[string]int
}

func NewTokenApp() *TokenApp{
	return &TokenApp{Accounts: map[string]int{}}
}

func (app *TokenApp) transfer(from,to crypto.Address,value int)error{
	if app.Accounts[from.String()] < value{
		return errors.New("not enough balance")
	}
	app.Accounts[from.String()] -= value
	app.Accounts[to.String()] += value
	return nil
}

func (app *TokenApp) issue(issuer,to crypto.Address,value int) error{
	if !bytes.Equal(issuer,SYSTEM_ISSUER){
		return errors.New("invalid issuer")
	}
	app.Accounts[to.String()] += value
	return nil
}

func (app *TokenApp)Dump(){
	fmt.Printf("state => %v\n",app.Accounts)
}

func main() {
	app := NewTokenApp()
	p1 := kf.GenPrivKey()
	p2 := kf.GenPrivKey()
	//app.issue(SYSTEM_ISSUER,p1.PubKey().Address(),1000)
	//app.transfer(p1.PubKey().Address(),p2.PubKey().Address(),50)
	app.Dump()

	txIssue := lib.NewTransaction(lib.NewIssuePayload(SYSTEM_ISSUER,p1.PubKey().Address(),big.NewInt(150)))
	txIssue.Sign(issuer)
	fmt.Printf("issue tx => %v \n",txIssue)
	fmt.Printf("validated => %t\n",txIssue.Verify())

	txTransfer := lib.NewTransaction(lib.NewTxPayload(p1.PubKey().Address(),p2.PubKey().Address(),big.NewInt(25)))
	txTransfer.Sign(p1)
	fmt.Printf("transfer tx => %v\n",txTransfer)
	fmt.Printf("validated => %t\n ",txTransfer.Verify())

}


