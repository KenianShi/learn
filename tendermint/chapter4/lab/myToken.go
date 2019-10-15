package lab
//
//import (
//	"bytes"
//	"errors"
//	"fmt"
//	"github.com/tendermint/tendermint/abci/server"
//	"github.com/tendermint/tendermint/abci/types"
//	"github.com/tendermint/tendermint/crypto"
//)
//
//var (
//	SYSTEM_ISSUER = crypto.Address("KING_OF_TOKEN")
//)
//
//type TokenApp struct {
//	types.BaseApplication
//	Accounts map[string]int64
//}
//
//func NewTokenApp() *TokenApp {
//	return &TokenApp{Accounts: map[string]int64{}}
//}
//
//func (app *TokenApp) issue(issuer, to crypto.Address, value int64) error {
//	if !bytes.Equal([]byte(issuer), SYSTEM_ISSUER) {
//		return errors.New("invalid issuer")
//	}
//	app.Accounts[to.String()] += value
//	return nil
//}
//
//func (app *TokenApp) transfer(from, to crypto.Address, value int64) error {
//	if app.Accounts[from.String()] < value {
//		return errors.New("not enough money")
//	}
//	app.Accounts[from.String()] -= value
//	app.Accounts[to.String()] += value
//	return nil
//}
//
//func (app *TokenApp) DeliverTx(req types.RequestDeliverTx) (rsp types.ResponseDeliverTx) {
//	switch req.Tx[0] {
//	case 0x01:
//		app.issue(SYSTEM_ISSUER, crypto.Address("1111"), 1000)
//	case 0x02:
//		app.transfer(crypto.Address("1111"), crypto.Address("2222"), 100)
//	default:
//		return types.ResponseDeliverTx{Code: 0, Log: "Bad tx"}
//	}
//	return
//}
//
////func (app *TokenApp) Query(req types.RequestQuery) (rsp types.ResponseQuery) {
////	address := crypto.Address(req.Data)
////	rsp.Key = req.Data
////	rsp.Value = codec.marshal
////
////}
//
//func (app *TokenApp) Dump() {
//	fmt.Printf("state ==> %v \n", app.Accounts)
//}
//
//func main() {
//	app := NewTokenApp()
//	svr, err := server.NewServer(":26658", "socket", app)
//	if err != nil {
//		panic(err)
//	}
//	svr.Start()
//	defer svr.Stop()
//	fmt.Println("app started")
//	select {}
//}
