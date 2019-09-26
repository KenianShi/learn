package chapter3

import (
	"os"
	"fmt"
	"strconv"
	"flag"
)

type CLI struct {
	BC *Blockchain
}


func (cli *CLI) printUsage(){
	fmt.Println("Usage:")
	fmt.Println("addBlock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs(){
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string){
	cli.BC.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) printChain(){
	bci := cli.BC.Iterator()
	for {
		block := bci.Next()
		fmt.Printf("PreHash:%x \nData: %s \nHash: %x \nNonce: %d \n",block.Prehash,block.Data,block.Hash,block.Nonce)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s \n",strconv.FormatBool(pow.Validate()))
		fmt.Println()
		if len(block.Prehash) == 0{
			break
		}
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()
	addBlockCmd := flag.NewFlagSet("addblock",flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain",flag.ExitOnError)
	addBlockData := addBlockCmd.String("data","","Block data")
	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {

		}
	case "printChain" :
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {

		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed(){
		if *addBlockData == ""{
			addBlockCmd.Usage()
			os.Exit(1)
		}else{
			cli.addBlock(*addBlockData)
		}
	}

	if printChainCmd.Parsed(){
		cli.printChain()
	}
}












