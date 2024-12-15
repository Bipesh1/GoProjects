package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Player represents a player in the game
type Player struct {
	ID       string
	Conn     *websocket.Conn
	IsDrawer bool
}

// Message represents a generic message structure
type Message struct {
	Type    string `json:"type"`    // Type of message ("drawing", "guess", "word")
	Content string `json:"content"` // Content of the message
	Sender  string `json:"sender"`  // Sender's identifier
	X0      int    `json:"x0,omitempty"`
	Y0      int    `json:"y0,omitempty"`
	X1      int    `json:"x1,omitempty"`
	Y1      int    `json:"y1,omitempty"`
	Color   string `json:"color,omitempty"`
}

var (
	drawingBroadcast = make(chan Message)
	playersMutex     = sync.Mutex{}
	upgrader         = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins for prototype purposes
		},
	}
	Players          = make(map[*Player]bool)
	currentWordIndex = 0
	gameStarted      = false
)

func handleConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket connected")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	player := &Player{
		ID:   r.RemoteAddr,
		Conn: ws,
	}

	// Add the player
	playersMutex.Lock()
	fmt.Println("Locked playersMutex in handleConnection (adding player)")
	Players[player] = true
	if len(Players) >= 2 && !gameStarted {
		fmt.Println("Enough players to start game. Game has not started yet.")
		gameStarted = true // Mark the game as started
		go startGame()
	}
	playersMutex.Unlock()
	fmt.Println("Unlocked playersMutex in handleConnection (after adding player)")

	defer func() {
		playersMutex.Lock()
		fmt.Println("Locked playersMutex in handleConnection (removing player)")
		delete(Players, player) // Remove player on disconnect
		playersMutex.Unlock()
		fmt.Println("Unlocked playersMutex in handleConnection (after removing player)")
		ws.Close() // Close the connection
	}()

	for {
		var msg Message
		if err := ws.ReadJSON(&msg); err != nil {
			log.Printf("Error: %v", err)
			break
		}

		// Handle incoming messages based on type
		switch msg.Type {
		case "drawing":
			// Forward drawing message to broadcast channel
			fmt.Println("Received drawing message")
			drawingBroadcast <- msg
		case "guess":
			// Handle guess message
			fmt.Printf("Player %s guessed: %s\n", msg.Sender, msg.Content)
		}
	}
}

func startGame() {
	fmt.Println("Inside startGame function")
	playersMutex.Lock()
	fmt.Println("Locked playersMutex in startGame")

	if len(Players) < 2 {
		fmt.Println("Not enough players to get started with")
		gameStarted = false   // Reset flag if not enough players
		playersMutex.Unlock() // Unlock only once here before returning
		return
	}

	playersMutex.Unlock()
	fmt.Println("Unlocked playersMutex in startGame")

	fmt.Println("Calling updateDrawer from startGame")
	updateDrawer()

	fmt.Println("Calling setWord from startGame")
	setWord()
}

func handleDrawingMessages() {
	for {
		msg := <-drawingBroadcast
		fmt.Println("Received a drawing message to broadcast")
		playersMutex.Lock()
		fmt.Println("Locked playersMutex in handleDrawingMessages")
		for player := range Players {
			// Broadcast the drawing message to all players
			if err := player.Conn.WriteJSON(msg); err != nil {
				log.Printf("Error: %v", err)
				player.Conn.Close()
				delete(Players, player)
			}
		}
		playersMutex.Unlock()
		fmt.Println("Unlocked playersMutex in handleDrawingMessages")
	}
}

func handleWordBroadcast(word string) {
	fmt.Println("Handle word broadcast")
	playersMutex.Lock()
	fmt.Println("Locked playersMutex in handleWordBroadcast")
	defer playersMutex.Unlock()

	for player := range Players {
		if player.IsDrawer {
			// Send the word to the drawer player
			wordMessage := Message{
				Type:    "word",
				Content: word,
				Sender:  "Server",
			}
			if err := player.Conn.WriteJSON(wordMessage); err != nil {
				log.Printf("Error sending word to drawer: %v", err)
			} else {
				log.Printf("Word '%s' sent to drawer: %s", word, player.ID)
			}
		}
	}
	fmt.Println("Unlocked playersMutex in handleWordBroadcast (deferred)")
}

func setWord() {
	fmt.Println("Inside setWord function")
	playersMutex.Lock()
	fmt.Println("Locked playersMutex in setWord")

	words := []string{"chrome", "cat", "house", "dog", "orange"}
	currentWord := words[currentWordIndex]
	
	// Unlock before calling handleWordBroadcast to avoid deadlock
	playersMutex.Unlock()
	fmt.Println("Unlocked playersMutex in setWord before broadcasting word")

	handleWordBroadcast(currentWord)

	// Increment currentWordIndex after broadcasting
	playersMutex.Lock()
	currentWordIndex++
	if currentWordIndex >= len(words) {
		currentWordIndex = 0 // Loop back to the beginning
	}
	fmt.Println("Unlocked playersMutex in setWord after broadcasting word")
	playersMutex.Unlock()
}


func updateDrawer() {
	fmt.Println("Inside updateDrawer function")
	playersMutex.Lock()
	fmt.Println("Locked playersMutex in updateDrawer")
	defer playersMutex.Unlock()

	var nextDrawer *Player
	var foundDrawer = false

	for player := range Players {
		if player.IsDrawer {
			player.IsDrawer = false
			foundDrawer = true
		} else if foundDrawer && nextDrawer == nil {
			nextDrawer = player
		}
	}
	if nextDrawer == nil {
		for player := range Players {
			player.IsDrawer = true
			break
		}
	}
	if nextDrawer != nil {
		nextDrawer.IsDrawer = true
	}
	fmt.Println("Unlocked playersMutex in updateDrawer (deferred)")
	fmt.Println("Drawer has been updated")
}

func main() {
	var r = mux.NewRouter()
	fs := http.FileServer(http.Dir("./static"))
	r.HandleFunc("/ws/", handleConnection)
	r.PathPrefix("/").Handler(fs)
	go handleDrawingMessages()
	fmt.Println("Server started at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Printf("Error: %v", err)
	}
}
