package main

import (
	"errors"
	"fmt"
	"github.com/KenianShi/learn/mytendermint/lib"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/rpc/client"
	"math/big"
)

var (
	cli = client.NewHTTP("http://localhost:26657","/websocket")
)

//daemon中转账以及issuer写死了，以后可以改进
//daemon中转账以及issuer写死了，以后可以改进
func main() {
	rootCmd := cobra.Command{
		Use:"cli",
	}

	walletCmd := &cobra.Command{Use:"init-wallet"}
	issueCmd := &cobra.Command{
		Use:"issue-tx",
		Run: func(cmd *cobra.Command, args []string) {
			issue()
		},
	}

	transferCmd := &cobra.Command{Use:"transfer-tx",
		Run: func(cmd *cobra.Command, args []string) {
			transfer()
		},
	}

	queryCmd := &cobra.Command{
		Use:"query",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("please tell me what you want to query")
			}
			name := args[1]
			query(name)
			return nil
		},
	}

	rootCmd.AddCommand(walletCmd)
	rootCmd.AddCommand(issueCmd)

	rootCmd.AddCommand(transferCmd)
	rootCmd.AddCommand(queryCmd)
	rootCmd.Execute()
	initWallet()

}

func issue(){
	wallet := lib.LoadWallet()
	tx := lib.NewTransaction(lib.NewIssuePayload(wallet.GetAddress("issuer"),wallet.GetAddress("Alice"),big.NewInt(20)))
	tx.Sign(wallet.GetPrivKey("issuer"))
	bz,err := lib.MarshalJSON(tx)
	if err != nil {
		panic(err)
	}
	ret,err := cli.BroadcastTxCommit(bz)
	if err != nil {
		panic(err)
	}
	fmt.Printf("issue tx ret => %+v\n",ret)
}

func transfer(){
	wallet := lib.LoadWallet()
	tx := lib.NewTransaction(lib.NewTxPayload(wallet.GetAddress("Alice"),wallet.GetAddress("Bob"),big.NewInt(10)))
	tx.Sign(wallet.GetPrivKey("Alice"))
	bz,err := lib.MarshalJSON(tx)
	if err != nil {
		panic(err)
	}
	ret,err := cli.BroadcastTxCommit(bz)
	if err != nil{
		panic(err)
	}
	fmt.Printf("transfer tx ret => %+v\n",ret)
}

func query(name string){
	wallet := lib.LoadWallet()
	ret,err := cli.ABCIQuery("",wallet.GetAddress(name))
	if err != nil {
		panic(err)
	}
	fmt.Printf("ret => %+v\n",ret)
}

func initWallet(){
	wallet := lib.NewWallet()
	wallet.GenPrivKey("issuer")
	wallet.GenPrivKey("Alice")
	wallet.GenPrivKey("Bob")
	wallet.GenPrivKey("Coco")
	wallet.Save()
}