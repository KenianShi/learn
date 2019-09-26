package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

var addr1 = flag.String("addr", "localhost:8080", "server ip address")

func main() {
	flag.Parse()
	log.SetFlags(3)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{Scheme: "ws", Host: *addr1, Path: "/echo"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("dail err: ", err)
		return
	}
	defer c.Close()
	done := make(chan struct{}, 1)
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read message err: ", err)
				return
			}
			log.Println("client recv message: ", string(message))
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			fmt.Println("执行done，退出")
			return
		case t := <-ticker.C:
			err = c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write message err: ", err)
				return
			}
		case <-interrupt:
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write message err: ", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			fmt.Println("interrupt  over!")
		}
	}
}

//package main
//
//import (
//	"flag"
//	"github.com/gorilla/websocket"
//	"log"
//	"net/url"
//	"os"
//	"os/signal"
//	"time"
//	"fmt"
//)
//
//var addr1 = flag.String("addr", "localhost:8080", "server http address")
//
//func main() {
//	flag.Parse()
//	log.SetFlags(3)
//	interrupt := make(chan os.Signal, 1)
//	signal.Notify(interrupt, os.Interrupt)
//
//	u := url.URL{Scheme: "ws", Host: *addr1, Path: "/echo"}
//	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
//	if err != nil {
//		log.Println("dail err: ", err)
//		return
//	}
//	defer c.Close()
//	done := make(chan struct{}, 1)
//	go func() {
//		defer close(done)
//		for {
//			_, mess, err := c.ReadMessage()
//			if err != nil {
//				log.Println("read message err: ", err)
//				break
//			}
//			log.Printf("client recv message: %s \n", mess)
//		}
//	}()
//
//	ticker := time.NewTicker(time.Second)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-done:
//			return
//		case t := <-ticker.C:
//			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
//			if err != nil {
//				log.Println("write message err: ", err)
//				break
//			}
//		case <-interrupt:
//			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
//			if err != nil {
//				log.Println("close websocket err: ", err)
//				return
//			}
//			select {
//			case <-done:
//			case <-time.After(time.Second):
//			}
//			fmt.Println("顺利推出")
//			return
//		}
//	}
//
//}
