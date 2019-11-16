package lib

import (
	"github.com/tendermint/tendermint/crypto"
	kf "github.com/tendermint/tendermint/crypto/secp256k1"
	"io/ioutil"
)

type Wallet struct {
	Keys 	map[string]crypto.PrivKey
}

func NewWallet() *Wallet{
	return &Wallet{Keys: map[string]crypto.PrivKey{}}
}

func (wallet *Wallet)Save(){
	bz,err := codec.MarshalJSON(wallet)
	if err != nil { panic(err) }
	ioutil.WriteFile("./wallet",bz,0644)
}

func LoadWallet()*Wallet{
	var w Wallet
	bz,err := ioutil.ReadFile("./wallet")
	if err != nil { panic(err) }
	err = codec.UnmarshalJSON(bz,&w)
	if err != nil {
		panic(err)
	}
	return &w
}

func (wallet *Wallet) GenPrivKey(name string) crypto.PrivKey{
	if privKey,exits := wallet.Keys[name];exits {
		return privKey
	}
	priv := kf.GenPrivKey()
	wallet.Keys[name] = priv
	return priv
}

func (wallet *Wallet) GetPrivKey(name string) crypto.PrivKey{
	return wallet.Keys[name]
}

func (wallet *Wallet) GetPubKey(name string) crypto.PubKey{
	if priv,exits := wallet.Keys[name];!exits{
		panic("key not exits")
	}else{
		return priv.PubKey()
	}
}

func (wallet *Wallet) GetAddress(name string) crypto.Address{
	if priv,exits := wallet.Keys[name];!exits{
		panic("key not exits")
	}else{
		return priv.PubKey().Address()
	}
}

