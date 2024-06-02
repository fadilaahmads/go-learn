package routes

import (
	"gotodolist/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine){
  router.GET("/", controllers.GetHomepage)
  router.GET("/todos", controllers.GetTodos)
 }
