package player

import (
	// "fmt"
	// "log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Player struct {
	ID   string
	Conn *websocket.Conn
}

type Message struct {
	Content string
	Sender  string
}

var Players = make(map[*Player]bool)
var Broadcast = make(chan Message)
var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}
