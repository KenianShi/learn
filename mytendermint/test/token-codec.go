package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/KenianShi/learn/mytendermint/lib"
	"github.com/tendermint/tendermint/crypto"
	kf "github.com/tendermint/tendermint/crypto/secp256k1"
)

var (
	issuer1 = kf.GenPrivKey()
	SYSTEM_ISSUER2 = issuer1.PubKey().Address()

)

type TokenApp2 struct {
	Accounts 	map[string]int
}

func NewTokenApp2()*TokenApp2{
	return &TokenApp2{Accounts: map[string]int{}}
}

//todo 此处使用codec.MarshalBinary，解析出来verify失败？原因待查明
func main() {
	p1 := kf.GenPrivKey()
	txIssue := lib.NewTransaction( lib.NewIssuePayload(
		issuer1.PubKey().Address(),
		p1.PubKey().Address(),
		1000))

	txIssue.Sign(issuer1)
	fmt.Printf("开始：%s \n",txIssue.PubKey.Address().String())
	fmt.Printf("validated1 => %t\n",txIssue.Verify())
	fmt.Printf("%T,%+v\n",txIssue,txIssue)
	rawtx,err := lib.MarshalBinary(txIssue)
	if err !=nil { panic(err) }


	var txReceived lib.Tx
	err =  lib.UnmarshalBinary(rawtx,&txReceived)
	if err != nil { panic(err) }
	fmt.Printf("开始：%s \n",txReceived.PubKey.Address().String())
	fmt.Printf("%T,%+v\n",txReceived,txReceived)
	fmt.Printf("validated2 => %t\n",txReceived.Verify())


	//
	//txIssue := lib.NewTransaction(lib.NewIssuePayload(issuer1.PubKey().Address(),p1.PubKey().Address(),big.NewInt(1000)))
	//txIssue.Sign(issuer1)
	//txSend,err := lib.MarshalBinary(txIssue)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("issue tx:\n%+v\n",txIssue)
	//
	//var txReceived lib.Tx
	//err = lib.UnmarshalBinary(txSend,&txReceived)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("\n\nissue tx decoded => \n%+v\n",txReceived)
	//fmt.Printf("validated => %t\n",txReceived.Verify())
	//
	////fmt.Printf("validated => %t\n",txIssue.Verify())
}


func (app *TokenApp2) Issue(issuer,to crypto.Address,value int)error{
	if !bytes.Equal(issuer,SYSTEM_ISSUER2){
		return errors.New("invalid issuer")
	}
	app.Accounts[to.String()] += value
	return nil
}

func (app *TokenApp2) Transfer(from,to crypto.Address,value int)error{
	if app.Accounts[from.String()] < value {
		return errors.New("not enough balance")
	}
	app.Accounts[from.String()] -= value
	app.Accounts[to.String()] += value
	return nil
}

func (app *TokenApp2) Dump(){
	fmt.Printf("state => %+v\n",app.Accounts)
}




