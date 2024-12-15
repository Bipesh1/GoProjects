package router

import(
	"github.com/gin-gonic/gin"
	"todo/controller"
)

func Registerrouters(router *gin.Engine){
router.GET("/todo", controller.Gettodo)
router.POST("/todo",controller.Createtodo)
}