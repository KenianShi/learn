package lab

import (
	"encoding/json"
	"github.com/tendermint/tendermint/crypto"
	kf "github.com/tendermint/tendermint/crypto/secp256k1"
	"io/ioutil"
)

type Wallet struct {
	Keys 		map[string]crypto.PrivKey
}

func NewWallet()*Wallet{
	return &Wallet{Keys: map[string]crypto.PrivKey{}}
}

func(w *Wallet) GenPriKey(label string)crypto.PrivKey{
	priv := kf.GenPrivKey()
	w.Keys[label] = priv
	return priv
}

func(w *Wallet) GetPrivKey(label string)crypto.PrivKey{
	return w.Keys[label]
}

func (w *Wallet) GetPubKey(label string) crypto.PubKey{
	priv,exits := w.Keys[label]
	if !exits{
		panic("keys not found")
	}
	return priv.PubKey()
}

func (w *Wallet) GetAddress(label string) crypto.Address{
	priv,exits := w.Keys[label]
	if !exits{
		panic("keys not found")
	}
	return priv.PubKey().Address()
}

func (w *Wallet) Save(path string) error{
	bz,err := json.Marshal(w)
	if err != nil {
		panic(err)
	}
	return ioutil.WriteFile(path,bz,0644)
}

func LoadWallet(path string)*Wallet{
	var w Wallet
	bz,err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bz,&w)
	if err != nil {
		panic(err)
	}
	return &w
}

