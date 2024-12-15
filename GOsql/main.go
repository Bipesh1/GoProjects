package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db,err:=sql.Open("mysql","root:admin@tcp(127.0.0.1:3306)/")
	if err!=nil{
		log.Fatal("Cannot create the database object")
	}
	router:= gin.Default()

	router.POST("createdatabase", func(c *gin.Context){
		dbname:= c.Query("name")
		if dbname==""{
			c.JSON(http.StatusBadRequest,gin.H{
				"message":"Database name must be provided",
			})
		}
		query:=fmt.Sprintf("CREATE DATABASE %s",dbname)
		_,err =db.Exec(query)
		if err!=nil{
			c.JSON(500,gin.H{
				"message":"Cannot create the database",
			})
		}else{
			c.JSON(200,gin.H{
				"message":"Database Created Successfully",
			})
		}
	})
	router.POST("createtable",func(c *gin.Context){
		tablename:= c.Query("name")
		if tablename==""{
			c.JSON(500,gin.H{
				"message":"Table name must be provided",
			})
		}
	})

	router.Run("0.0.0.0:9000")
	

}