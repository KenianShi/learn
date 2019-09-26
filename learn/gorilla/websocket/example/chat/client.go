package chat

import (
	"github.com/gorilla/websocket"
	"time"
	"log"
	"bytes"
)

const (
	writeWait		= time.Second * 10
	pongWait		= time.Second * 10
	pingPeriod		= (pongWait * 9)/10
	maxMessageSize 	= 512
)

var (
	newLine		= []byte{'\n'}
	space 		= []byte{' '}
)


type Client struct {
	hub		*Hub
	conn 	*websocket.Conn

	//Buffered channel of outbound message
	send 	chan []byte
}

func (c *Client) readPump(){
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pingPeriod))
		return nil
	})

	for {
		_,message,err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err,websocket.CloseGoingAway,websocket.CloseAbnormalClosure){
				log.Printf("error: %v",err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message,newLine,space,-1))
		c.hub.broadcast <- message
	}
}

func (c *Client) writePump(){
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := c.send:




		}






	}



}




