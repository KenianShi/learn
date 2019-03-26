package main

import "shikenian/learn/blockchain/sknblockchain/chapter3"

func main() {
	bc := chapter3.NewBlockchain()
	defer bc.Db.Close()
	cli := chapter3.CLI{bc}
	cli.Run()

}
