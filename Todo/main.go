package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	// "os"
	// "os/signal"
	// "strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thedevsaddam/renderer"
	// mgo "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

var rnd *renderer.Render
var db *mongo.Database

const (
	hostName       string = "mongodb+srv://bipesh:CSIvZZ2uqs4o5Vtf@cluster0.zd9twdf.mongodb.net/"
	dbName         string = "demo_todo"
	collectionName string = "todo"
	port           string = ":9000"
)

type (
	todoModel struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		Title     string             `bson:"title"`
		Completed bool               `bson:"completed"`
		CreatedAt time.Time          `bson:"createdAt"`
	}
	todo struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		CreatedAt time.Time `json:"createdat"`
	}
)

func init() {
	rnd = renderer.New()
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(hostName))
	checkErr(err)
	err = client.Ping(context.Background(), nil)
	db = client.Database(dbName)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := rnd.Template(w, http.StatusOK, []string{"static/home.tpl"}, nil)
	checkErr(err)
}
func fetchTodos(w http.ResponseWriter, r *http.Request) {
	todosdata := []todo{}
	cursor, err := db.Collection(collectionName).Find(context.Background(), bson.M{})
	if err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to fetch todo data",
		})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var t todoModel
		if err := cursor.Decode(&t); err != nil {
			checkErr(err)
		}
		todosdata = append(todosdata, todo{
			ID:        t.ID.Hex(),
			Title:     t.Title,
			Completed: t.Completed,
			CreatedAt: t.CreatedAt,
		})
	}
	rnd.JSON(w, http.StatusOK, renderer.M{
		"todosdata": todosdata,
	})
}
func createTodo(w http.ResponseWriter, r *http.Request) {
	tododata := todo{}
	err := json.NewDecoder(r.Body).Decode(&tododata)
	checkErr(err)
	tododata_db := todoModel{
		ID:        primitive.NewObjectID(),
		Title:     tododata.Title,
		Completed: tododata.Completed,
		CreatedAt: time.Now(),
	}
	_, err = db.Collection(collectionName).InsertOne(context.Background(), tododata_db)
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, renderer.M{
			"message": "Failed to insert todo data",
		})
		return
	}

}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Mount("/todo", todoHandlers())

	srv := http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("Listening on port", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen:%s\n", err)
		}
	}()
	log.Println("Server should be running at this point...")
	select{}
}
func todoHandlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", fetchTodos)
		r.Post("/", createTodo)
		// r.Put("/{id}",updateTodo)
		// r.Delete("/{id}",deleteTodo)
	})
	return rg
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
