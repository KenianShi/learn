package main

import (
	"fmt"
	"github.com/KenianShi/learn/mytendermint/lib"
	"github.com/tendermint/tendermint/abci/server"
)

func main() {
	app := lib.NewTokenApp()
	svr,err := server.NewServer(":26658","socket",app)
	if err != nil {
		panic(err)
	}
	svr.Start()
	defer svr.Stop()
	fmt.Printf("token server started...")
	select {}
}