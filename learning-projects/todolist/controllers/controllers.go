package controllers

import (
	"gotodolist/services"
	"gotodolist/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetHomepage(c *gin.Context){
  c.JSON(http.StatusOK, gin.H{"message":"Welcome to the To-Do-List App!"}) 
}

func GetTodos(c *gin.Context){
  todos, err := services.GetTodos()
  if err != nil{
    utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
  }
  utils.ResponseWithJSON(c, http.StatusOK, todos)
}

func GetTodoById(c *gin.Context){
  todoId := c.Param("id")
  intTodoId, err := strconv.Atoi(todoId)
  if err != nil{
    log.Panicln("Failed to convert ID")
    utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
  }
  todo, err := services.GetTodosById(intTodoId)
  if err != nil {
    utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
  }
  utils.ResponseWithJSON(c, http.StatusOK, todo)
}
