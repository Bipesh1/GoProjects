package main

import(
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"log"
)

type Book struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Price string `json:"price"`
}
var books []Book
var i int=1

func fetchAllBooks(w http.ResponseWriter, r *http.Request){

}
func fetchBookById(w http.ResponseWriter,r *http.Request){

}
func CreateBook(w http.ResponseWriter,r *http.Request){
b:= Book{}
if err:= json.NewDecoder(r.Body).Decode(&b);err!=nil{
	log.Fatal(err)
}
b.ID= i
books=append(books, b)
json.NewEncoder(w).Encode(books)
}

func main(){
r:= chi.NewRouter()

r.Get("/books",fetchAllBooks)
r.Get("/books/{id}",fetchBookById)
r.Post("/books/", CreateBook)

if err:= http.ListenAndServe(":8080",r);err!=nil{
	log.Fatal(err)
}
}