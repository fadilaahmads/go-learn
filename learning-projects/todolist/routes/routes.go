package routes

import (
	"gotodolist/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine){
  router.GET("/", controllers.GetHomepage)
  router.GET("/todos", controllers.GetTodos)
  router.GET("/todos/:id", controllers.GetTodoById)
  router.POST("/todos", controllers.CreateTodoHandler)
  router.POST("/todos/:id", controllers.UpdateTodoHandler)
  router.DELETE("/todos/:id", controllers.DeleteTodoHandler)
 }
