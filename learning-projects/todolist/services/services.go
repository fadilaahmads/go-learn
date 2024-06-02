package services

import (
	"fmt"
	"gotodolist/databases"
	"gotodolist/models"
	"log"
)

func GetTodos()([]models.Todo, error){
  var todos []models.Todo
  db := databases.GetDB()
  if db == nil {
    log.Println("Database connection is not initialize")
    return nil, fmt.Errorf("Database is not initialize")
  }

  // querying data 
  if err := db.Find(&todos).Error; err != nil{
    log.Printf("Error receiving todos: %v", err)
    return nil, err
  }

  return todos,nil
}
