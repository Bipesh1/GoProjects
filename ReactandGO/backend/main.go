package main

import (
	"log"
	 "todo/router"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	r:= gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Allowed HTTP methods
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
        ExposeHeaders:    []string{"Content-Length", "X-Custom-Header"}, // Expose headers
        AllowCredentials: true, // Allow credentials (cookies, authorization headers, etc.)
        MaxAge:           12 * time.Hour, // Cache preflight response for 12 hours
	}))
	router.Registerrouters(r)

	if err:=r.Run("0.0.0.0:9000");err!=nil{
		log.Printf("Error: %v",err)
	}

	

}