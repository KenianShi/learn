package main

import (
	"encoding/json"
	"fmt"
	kf "github.com/tendermint/tendermint/crypto/ed25519"
)


type Letter struct {
	Msg 		[]byte
	Signature 	[]byte
	PubKey 		kf.PubKeyEd25519
}

func main() {
	byte := sign("Hello my verify")
	msg := verify(byte)
	fmt.Println(msg)

}

func verify(letterByte []byte)string{
	letter := Letter{}
	err := json.Unmarshal(letterByte,&letter)
	if err != nil {
		fmt.Printf("解析失败，err：%s \n",err)
		panic(err)
	}
	pub := letter.PubKey
	signature := letter.Signature
	msg := letter.Msg
	if pub.VerifyBytes(msg,signature){
		return string(msg)
	}else{
		fmt.Printf("验证不通过 \n")
		return "验证不通过"
	}
}

func sign(msg string)[]byte{
	priv := kf.GenPrivKey()
	pub := priv.PubKey()

	signature,err := priv.Sign([]byte(msg))
	if err != nil {
		panic(err)
	}
	letter := &Letter{
		Msg:       []byte(msg),
		Signature: signature,
		PubKey:    pub.(kf.PubKeyEd25519),
	}

	byte,err := json.Marshal(letter)
	if err != nil {
		panic(err)
	}
	return byte
}