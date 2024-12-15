package routes

import (
	// "net/http"
	"fmt"
	"gosimplerealchat/controller"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(r *mux.Router) {
	fmt.Println("it is here")
	r.HandleFunc("/ws/", controller.HandleConnections)
}
