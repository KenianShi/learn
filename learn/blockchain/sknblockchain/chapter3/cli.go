package chapter3

import (
	"fmt"
	"os"
	"strconv"
	"flag"
)

type CLI struct {
	Bc		*Blockchain
}

func (cli *CLI) printUsage(){
	fmt.Println("Usage")
	fmt.Println("addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("printlnchain - print all the blocks of blockchain")
}

func (cli *CLI) validateArgs(){
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string){
	cli.Bc.AddBlock(data)
	fmt.Println("Success")
}

func (cli *CLI) printChain(){
	fmt.Println("go into printchain")
	bci := cli.Bc.Iterator()
	for{
		block := bci.Next()
		fmt.Printf("PrevHash:%x Data:%s Hash:%x ",block.PrevBlockHash,block.Data,block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s \n",strconv.FormatBool(pow.Validate()))
		fmt.Println()
		if len(block.PrevBlockHash) == 0{
			break
		}
	}
}

func (cli *CLI) Run(){
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock",flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain",flag.ExitOnError)
	addBlockData := addBlockCmd.String("data","","Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		checkErr(err)
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		checkErr(err)
	default:
		fmt.Println("cannot find the set")
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed(){
		if *addBlockData == ""{
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed(){
		cli.printChain()
	}

}


