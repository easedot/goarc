package router

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"

	"github.com/easedot/goarc/domain"
)

const (
	Channel = "ease_chat"
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

var (
	entering = make(chan client)
	leaving  = make(chan client)
	out      = make(chan clientMsg, 100)
	poll     = make(chan domain.Message)
	push     = make(chan domain.Message)
)

func fetchMessage(rds *redis.Pool, ch chan domain.Message) {
	conn := rds.Get()
	defer conn.Close()
	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(Channel)
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			var msg domain.Message
			json.Unmarshal(v.Data, &msg)
			ch <- msg
		}
	}
}

func writeMessage(rds *redis.Pool, ch chan domain.Message) {
	for input := range ch {
		msg, err := json.Marshal(input)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}
		writeToMsgCenter(rds, msg)
	}
}
func writeToMsgCenter(rds *redis.Pool, msg []byte) {
	conn := rds.Get()
	defer conn.Close()
	if err := conn.Send("PUBLISH", Channel, msg); err != nil {
		log.Printf("error: %v", err)
	}
	if err := conn.Flush(); err != nil {
		log.Printf("error: %v", err)
	}
}

func handleMessages() {
	clients := make(map[client]struct{}) // connected clients

	for {
		select {
		case msg := <-poll:
			for cli := range clients {
				select {
				case out <- clientMsg{Msg: msg, Conn: cli.Conn}:
				default:
				}
			}
		case cli := <-entering:
			clients[cli] = struct{}{}
		case cli := <-leaving:
			delete(clients, cli)
		}
	}
}

func NewChatServer(e *echo.Echo, rds *redis.Pool) {
	go handleMessages()
	go writeMessage(rds, push)
	go fetchMessage(rds, poll)
	go clientWrite(out)

	e.GET("/ws", func(c echo.Context) error {
		handleConnections(c, rds)
		return nil
	})
}
func handleConnections(c echo.Context, rds *redis.Pool) {
	var upgrader = websocket.Upgrader{}
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	who := c.RealIP()

	timer := time.NewTimer(Timeout)
	go func() {
		<-timer.C
		log.Print("Timeout 1 minute")
		ws.Close()
		timer.Stop()
	}()

	cli := client{Conn: ws, Name: c.RealIP()}
	entering <- cli
	push <- domain.Message{Message: who + " has arrived"}

	for {
		// Read in a new message as JSON and map it to a Message object
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		timer.Reset(Timeout)
		// Send the newly received message to the poll channel
		writeToMsgCenter(rds, msg)
	}
	leaving <- cli
	push <- domain.Message{Message: who + " has left"}
}

func clientWrite(clientChan <-chan clientMsg) {
	for cm := range clientChan {
		err := cm.Conn.WriteJSON(cm.Msg)
		if err != nil {
			log.Printf("error: %v", err)
		}
	}
}
