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

  if err := utils.ValidateID(intTodoId); err != nil {
    utils.ResponseWithError(c, http.StatusBadRequest, err.Error())
    return
  }
  todo, err := services.GetTodosById(intTodoId)
  if err != nil {
    utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
  }
  utils.ResponseWithJSON(c, http.StatusOK, todo)
}

func CreateTodoHandler(c *gin.Context){
  var todo models.Todo
  if err:= c.ShouldBindJSON(&todo); err != nil{
    utils.ResponseWithError(c, http.StatusBadRequest, "Invalid Input")
    return
  }

  sanitizedTodo, err := utils.ValidateSanitizedTodo(todo.Title)
  if err != nil {
    utils.ResponseWithError(c, http.StatusBadRequest, err.Error())
  }

  todo.Title = sanitizedTodo

  if err:= services.AddTodo(&todo); err != nil {
    utils.ResponseWithError(c, http.StatusInternalServerError, "Could not create todo")
    return
  }

  utils.ResponseWithJSON(c, http.StatusCreated, todo)
}

func UpdateTodoHandler(c *gin.Context){
  var todo models.Todo 
  // parsing request arguments to todo struct
  if err := c.ShouldBindJSON(&todo); err != nil{
    utils.ResponseWithError(c, http.StatusBadRequest, "Invalid input")
  }
  // id check 
  todoId, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    log.Println("failed to convert id to int: ", err)
    return
  }

  if err := utils.ValidateID(todoId); err != nil{
    utils.ResponseWithError(c, http.StatusBadRequest, err.Error())
    return
  } 
  todo.ID = todoId 
  
  // title check
  sanitizedTitle, err := utils.ValidateSanitizedTodo(c.Param("title"))
  if err != nil{
    utils.ResponseWithError(c, http.StatusBadRequest, err.Error())
  }
  todo.Title = sanitizedTitle

  completedToBoolean, err := strconv.ParseBool(c.Param("completed"))
  if err != nil{
    utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
  }
  todo.Completed = completedToBoolean
 
  if err := services.UpdateTodo(&todo.ID, &todo.Title, &todo.Completed); err != nil{
    utils.ResponseWithJSON(c, http.StatusInternalServerError, "Could not update todo")
  }

  utils.ResponseWithJSON(c, http.StatusOK, todo)
}

func DeleteTodoHandler(c *gin.Context){
  var todoId string = c.Param("id")
  intTodoId, err := strconv.Atoi(todoId)
  if err != nil {
    utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
  }

  if err := services.DeleteTodo(intTodoId); err != nil {
    utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
  }
  utils.ResponseWithJSON(c, http.StatusOK, "Todo successfully deleted")
}
