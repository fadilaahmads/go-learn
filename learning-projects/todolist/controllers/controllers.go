package controllers

import (
	"gotodolist/models"
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

func CreateTodoHandler(c gin.Context){
  var todo models.Todo
  if err:= c.ShouldBindJSON(&todo); err =! nil{
    utils.ResponseWithError(c, http.StatusBadRequest, "Invalid Input")
  }

  if err:= services.AddTodo(&todo); err != nil {
    utils.ResponseWithError(c, http.StatusInternalServerError, "Could not create todo")
    return
  }

  utils.ResponseWithJSON(c, http.StatusCreated, todo)
}
