package main

import (
	"fmt"
	"github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
)

//这个例子用来感受ABCI消息流程的顺序checkTx-->beginBlock-->deliverTx-->endBlock-->commit
type EzApp struct{
	types.BaseApplication
}

func NewEzApp() *EzApp{
	return &EzApp{}
}

func (app *EzApp) InitChain(req types.RequestInitChain)(rsp types.ResponseInitChain){
	fmt.Printf("初始化 initChain ==> %+v \n",req)
	return
}


func (app *EzApp) Info(req types.RequestInfo)(rsp types.ResponseInfo){
	fmt.Printf("信息 info ==> %+v \n",req)
	return
}

func (app *EzApp) Query(req types.RequestQuery) (rsp types.ResponseQuery){
	fmt.Printf("查询 query ==> %+v \n",req)
	return
}

func (app *EzApp) CheckTx(req types.RequestCheckTx)(rsp types.ResponseCheckTx){
	fmt.Printf("检查交易 CheckTx ==> %+v \n",req)
	return
}

func (app *EzApp) BeginBlock(req types.RequestBeginBlock)(rsp types.ResponseBeginBlock){
	fmt.Printf("开始区块 BeginBlock ==> %+v \n",req)
	return
}

func (app *EzApp) DeliverTx(req types.RequestDeliverTx)(rsp types.ResponseDeliverTx){
	fmt.Printf("传送交易 deliverTx ==> %+v \n",req)
	return
}

func (app *EzApp) EndBlock(req types.RequestEndBlock)(rsp types.ResponseEndBlock){
	fmt.Printf("结束区块 EndBlock ==> %+v \n",req)
	return
}

func (app *EzApp) Commit()(rsp types.ResponseCommit){
	fmt.Printf("提交区块 commit ==> \n")
	return
}

func main() {
	app := NewEzApp()
	svr,err := server.NewServer(":26658","socket",app)
	if err != nil {
		panic(err)
	}
	svr.Start()
	defer svr.Stop()
	fmt.Println("abci server started")
	select {}

}









