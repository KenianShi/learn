package main

import (
	"net/rpc"
	"net"
	"log"
	"time"
)

type Int int

type Args struct {
	A int `json:"a"`
	B int `json:"b"`
}

func (i *Int) Sum(args *Args,reply *int) error{
	*reply = args.A + args.B
	return nil
}

type MultyArgs struct{
	A *Args `json:"aa"`
	B *Args `json:"bb"`
}

type MultyReply struct {
	A int `json:"aa"`
	B int `json:"bb"`
}

func (i *Int) Multy(args *MultyArgs,reply *MultyReply) error{
	reply.A = args.A.A * args.A.B
	reply.B = args.B.A * args.B.B
	return nil
}

func main() {
	newServe := rpc.NewServer()
	i := new(Int)
	newServe.Register(i)
	lis,err := net.Listen("tcp","127.0.0.1:1234")
	if err != nil {
		log.Fatal("ERR:",err)
	}
	go newServe.Accept(lis)
	newServe.HandleHTTP("/foo","/bar")
	time.Sleep(time.Second*20000000000)


}
