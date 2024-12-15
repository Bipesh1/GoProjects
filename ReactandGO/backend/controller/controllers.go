package controller

import (
	// "fmt"
	"todo/models"

	"github.com/gin-gonic/gin"
)

var todos []models.Todo

func Gettodo(c *gin.Context) {
	var err error
	todos, err = models.Gettodos()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "An error occured while fetching data",
		})
	}
	if len(todos) == 0 {
		c.JSON(200, gin.H{
			"message": "No data to show",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "successful",
			"todos":   todos,
		})
	}
}
func Createtodo(c *gin.Context) {
t:= models.Todo{}
if err:= c.ShouldBindJSON(&t);err!=nil{
	c.JSON(500,gin.H{
		"message":"Cannot parse into json",
	})
}
err:=models.Createtodos(t)
if(err!=nil){
	c.JSON(500,gin.H{
		"message":"Cannot create the todo list",
	})
}else{
	c.JSON(200,gin.H{
		"status":"Successful",
		"message":"Todo created",
	})
}

}
