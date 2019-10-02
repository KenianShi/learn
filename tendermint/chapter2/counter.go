package main

import (
	"encoding/json"
	"fmt"
	"github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
	"io/ioutil"
)

type CounterApp struct {
	types.BaseApplication
	Value 	int
	Version int64
	Hash 	[]byte
	History map[int64]int
}

func NewCounterApp() *CounterApp{
	app := CounterApp{History: map[int64]int{}}
	state,err := ioutil.ReadFile("./counter.state")
	if err != nil {
		return &app
	}
	err = json.Unmarshal(state,&app)
	if err != nil {
		return &app
	}
	return &app
}


//必须req.Tx[0] < 0x04的交易才能通过。checkTx是最先执行的一步
func (app *CounterApp) CheckTx(req types.RequestCheckTx)(rsp types.ResponseCheckTx){
	if req.Tx[0] <0x04{
		rsp.Log = "tx accepted"
		return
	}
	rsp.Code = 1
	rsp.Log = "bad tx rejected"
	return
}

func (app *CounterApp) DeliverTx(req types.RequestDeliverTx)(rsp types.ResponseDeliverTx){
	switch req.Tx[0] {
	case 0x01: app.Value += 1
	case 0x02: app.Value -= 1
	case 0x03: app.Value = 0
	default: return types.ResponseDeliverTx{Code:0,Log:"bad tx"}
	}
	rsp.Log = fmt.Sprintf("state updated : %d",app.Value)
	return
}

func (app *CounterApp) InitChain(req types.RequestInitChain)(rsp types.ResponseInitChain){
	var state map[string]int
	json.Unmarshal(req.AppStateBytes,&state)
	app.Value = state["counter"]
	return types.ResponseInitChain{}
}

func (app *CounterApp) Query(req types.RequestQuery) types.ResponseQuery{
	height := req.Height
	if req.Height == 0 {
		height = app.Version
	}
	val := fmt.Sprintf("%d",app.History[height])
	return types.ResponseQuery{Value:[]byte(val),Log:val}
}

func (app *CounterApp) Commit() (rsp types.ResponseCommit){
	fmt.Println("执行了一次commit")       //用这种方法可以得出如果重启后，程序将会从第一一块开始便利所有的交易，系统会自动重放之前的交易
	app.Version += 1
	app.Hash = []byte(fmt.Sprintf("",app.Version))    //近似将version当作hash
	app.History[app.Version] = app.Value
	state,err := json.Marshal(app)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("./counter.,state",state,0644)
	return
}

func (app *CounterApp) Info(req types.RequestInfo)(rsp types.ResponseInfo){
	return types.ResponseInfo{LastBlockHeight:app.Version,LastBlockAppHash:app.Hash}

}

func main() {
	app := NewCounterApp()
	svr,err := server.NewServer(":26658","socket",app)
	if err != nil {
		panic(err)
	}
	svr.Start()
	defer svr.Stop()
	fmt.Println("abci server started")
	select {}
}