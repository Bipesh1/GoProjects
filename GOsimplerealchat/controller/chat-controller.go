package controller

import (
	"gosimplerealchat/pkg/player"
	"log"
	"net/http"
	"fmt"
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	fmt.Println("It came here")
	ws, err := player.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	playerinfo := &player.Player{
		ID:   r.RemoteAddr,
		Conn: ws,
	}
	player.Players[playerinfo] = true

	for {
		var msg player.Message
		if err := ws.ReadJSON(&msg); err != nil {
			log.Printf("Error: %v", err)
			delete(player.Players, playerinfo)
		}
		player.Broadcast <- msg
	}
}
func HandleMessages() {
	for {
		msg := <-player.Broadcast
		fmt.Println(msg)
		for playerinfo := range player.Players {
			if err := playerinfo.Conn.WriteJSON(msg); err != nil {
				log.Printf("Error: %v", err)
				playerinfo.Conn.Close()
				delete(player.Players, playerinfo)
			}
		}
	}
}
