package main

import (
	"fmt"
	"github.com/KenianShi/learn/mytendermint/lib"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{
		Use:"token-wallet",
	}

	initCmd := &cobra.Command{
		Use:"init",
		Short:"init wallet",
		Run:func(cmd *cobra.Command,args []string){initWallet()},
	}

	loadCmd := &cobra.Command{
		Use:"load",
		Short:"load wallet",
		Run: func(cmd *cobra.Command, args []string) {
			loadWallet()
		},
	}

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(loadCmd)
	rootCmd.Execute()
}


func initWallet(){
	wallet := lib.NewWallet()
	wallet.GenPrivKey("issuer")
	wallet.GenPrivKey("michal")
	wallet.GenPrivKey("bob")
	fmt.Printf("wallet => %v\n",wallet)
	wallet.Save()
}

func loadWallet(){
	wallet := lib.LoadWallet()
	fmt.Printf("wallet => %+v,\n",wallet)
	priv1 := wallet.GetPrivKey("michal")
	fmt.Printf("get michal privKey: %+v\n",priv1)
	pubKey2 := wallet.GetPubKey("bob")
	fmt.Printf("get bob pubKeys: %+v\n",pubKey2)
	addr3 := wallet.GetAddress("issuer")
	fmt.Printf("get issuer addr: %+v\n",addr3)
}
