package main

import "shikenian/learn/blockchain/chapter4"

func main() {
	bc := chapter4.NewBlockchain()
	defer bc.Db.Close()
	cli := &chapter4.CLI{bc}
	cli.Run()
}