package router

import (
	"encoding/json"
	"log"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/gomodule/redigo/redis"

	//"github.com/gorilla/websocket"
	//"github.com/labstack/echo/v4"
	"github.com/easedot/goarc/domain"
)

const (
	//Channel = "ease_chat"
	Timeout = 1 * time.Minute
)

type client struct {
	Name string
	Conn *websocket.Conn
}
type clientMsg struct {
	Msg  domain.Message
	Conn *websocket.Conn
}
type chat struct {
	channel  string
	entering chan client
	leaving  chan client
	out      chan clientMsg
	poll     chan domain.Message
	push     chan domain.Message
	rds      *redis.Pool
}

func (c *chat) pollFromCenter() {
	conn := c.rds.Get()
	defer conn.Close()
	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(c.channel)
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			var msg domain.Message
			json.Unmarshal(v.Data, &msg)
			c.poll <- msg
		}
	}
}

func (c *chat) pushToCenter() {
	for input := range c.push {
		msg, err := json.Marshal(input)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}
		c.writeToMsgCenter(msg)
	}
}
func (c *chat) writeToMsgCenter(msg []byte) {
	conn := c.rds.Get()
	defer conn.Close()
	if err := conn.Send("PUBLISH", c.channel, msg); err != nil {
		log.Printf("error: %v", err)
	}
	if err := conn.Flush(); err != nil {
		log.Printf("error: %v", err)
	}
}

func (c *chat) handleMessages() {
	clients := make(map[client]struct{}) // connected clients

	for {
		select {
		case msg := <-c.poll:
			for cli := range clients {
				select {
				case c.out <- clientMsg{Msg: msg, Conn: cli.Conn}:
				default:
				}
			}
		case cli := <-c.entering:
			clients[cli] = struct{}{}
		case cli := <-c.leaving:
			delete(clients, cli)
		}
	}
}
func (c *chat) clientWrite() {
	for cm := range c.out {
		err := cm.Conn.WriteJSON(cm.Msg)
		if err != nil {
			log.Printf("error: %v", err)
		}
	}
}

func NewChatServer(rds *redis.Pool) *chat {
	chat := chat{
		channel:  "ease_chat",
		entering: make(chan client),
		leaving:  make(chan client),
		out:      make(chan clientMsg, 100),
		poll:     make(chan domain.Message),
		push:     make(chan domain.Message),
		rds:      rds,
	}

	go chat.handleMessages()
	go chat.pushToCenter()
	go chat.pollFromCenter()
	go chat.clientWrite()
	return &chat
}

func (c *chat) ChatHandle(ws *websocket.Conn) {
	who := "jhh"
	cli := client{Conn: ws, Name: who}
	c.entering <- cli
	c.push <- domain.Message{Message: who + " has arrived"}

	for {
		// Read in a new message as JSON and map it to a Message object
		ws.SetReadDeadline(time.Now().Add(Timeout))
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		// Send the newly received message to the poll channel
		c.writeToMsgCenter(msg)
	}

	c.leaving <- cli
	c.push <- domain.Message{Message: who + " has left"}
}
