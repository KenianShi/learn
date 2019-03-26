package chapter4

import (
	"fmt"
	"os"
	"strconv"
	"flag"
)

type CLI struct {}

func (cli *CLI) creatBlockchain(address string){
	bc := CreateBlockchain(address)
	bc.db.Close()
	fmt.Println("Done!")
}

func (cli *CLI) getBalance(address string){
	bc := CreateBlockchain(address)
	defer bc.db.Close()
	balance := 0
	UTXOs := bc.FindUTXO(address)
	for _,out := range UTXOs{
		balance += out.Value
	}
	fmt.Printf("Balance of '%s':%d\n",address,balance)
}

func(cli *CLI) printUsage(){
	fmt.Println("Usage:")
	fmt.Println("getBalance -address ADDRESS - get balance of ADDRESS")
	fmt.Println("createBlockchain -address ADDRESS - Create a blockchain and send genesis block reward to Address")
	fmt.Println("printchain - Print all the blocks of blockchain")
	fmt.Println("send -from FROM -to TO -amount AMOUNT -Send Amount of coins from FROM address to TO")
}

func (cli *CLI)validateArgs(){
	if len(os.Args) < 2{
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) printChain(){
	bc := CreateBlockchain("")
	defer bc.db.Close()
	bci := bc.Iterator()
	for {
		block := bci.Next()
		fmt.Printf("Prev. hash:%x  Hash:%x   \n",block.PreBlockHash,block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s \n",strconv.FormatBool(pow.Validate()))
		fmt.Println()
		if len(block.PreBlockHash) == 0{
			break
		}
	}
}

func (cli *CLI) send(from,to string,amount int){
	bc := CreateBlockchain(from)
	defer bc.db.Close()
	tx := NewUTXOTransaction(from,to,amount,bc)
	bc.MineBlock([]*Transaction{tx})
	fmt.Println("Mined Success!")
}

func (cli *CLI) Run(){
	cli.validateArgs()

	getBalanceCmd := flag.NewFlagSet("getBalance",flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createBlockchain",flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send",flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain",flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address","","The address to get balance for")
	createBlockchainAddress := createBlockchainCmd.String("address","","The address to send genesis block reward to")
	sendFrom := sendCmd.String("from","","Source Wallet address")
	sendTo := sendCmd.String("to","","Destination wallet address")
	sendAmount := sendCmd.Int("amount",0,"Amount to send")

	switch os.Args[1]{
	case "getBalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		checkErr(err)

	case "createBlockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		checkErr(err)
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		checkErr(err)
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		checkErr(err)
	default:
		cli.printUsage()
		os.Exit(1)
	}
	if getBalanceCmd.Parsed(){
		if *getBalanceAddress == ""{
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress)
	}

	if createBlockchainCmd.Parsed(){
		if *createBlockchainAddress == ""{
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.creatBlockchain(*createBlockchainAddress)
	}

	if printChainCmd.Parsed(){
		cli.printChain()
	}

	if sendCmd.Parsed(){
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0{
			sendCmd.Usage()
			os.Exit(1)
		}
		cli.send(*sendFrom,*sendTo,*sendAmount)
	}


}























