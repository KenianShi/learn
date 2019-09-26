package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"net/url"
	"github.com/gorilla/websocket"
)

var clientAddr = flag.String("addr","127.0.0.1:8080","server ip address")

func main() {
	flag.Parse()
	log.SetFlags(3)
	interrupt := make(chan os.Signal,1)
	signal.Notify(interrupt,os.Interrupt)
	done := make(chan struct{},1)

	u := url.URL{Scheme:"ws",Host:*clientAddr,Path:"/echo"}
	c,_,err := websocket.DefaultDialer.Dial(u.String(),nil)
	if err != nil {
		log.Println("dail err: ",err)
		return
	}
	defer c.Close()

	go func(){
		defer close(done)
		for {





		}

	}()


}
