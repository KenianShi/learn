package main

import "shikenian/learn/blockchain/btc"

func main() {
	bc := btc.NewBlockchain()
	defer bc.Db.Close()
	cli := btc.CLI{bc}
	cli.Run()
}