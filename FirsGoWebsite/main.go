package main

import(
	"net/http"
	"html/template"
	"path/filepath"
	"fmt"
	
)

type PageData struct{
	Title,Content string
}

func renderTemplate(w http.ResponseWriter, tmpl string,data PageData){
	tmplPath:= filepath.Join("templates",tmpl)
	t,err:= template.ParseFiles(tmplPath)
	if(err!=nil){
		http.Error(w,"Unable to load template",http.StatusInternalServerError)
		return
	}
	t.Execute(w,data)


}

func homeHandler(w http.ResponseWriter, r *http.Request){
data:= PageData{Title: "Home", Content:"This is home"}
renderTemplate(w,"index.html",data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
data := PageData{Title: "About",Content: "This is an about page"}
renderTemplate(w,"about.html",data)
}


func main(){
	http.HandleFunc("/",homeHandler)
	http.HandleFunc("/about",aboutHandler)
	fmt.Println("Starting server at port 8080")
	if err:=http.ListenAndServe(":8080",nil);err!=nil{
		fmt.Println(err)
	}

}