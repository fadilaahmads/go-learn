package controllers

import (
	"gotodolist/services"
	"gotodolist/utils"
	"net/http"

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


