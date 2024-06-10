package services

import (
	"errors"
	"gotodolist/databases"
	"gotodolist/models"
	"log"

	/*   "github.com/jinzhu/gorm" */
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	db := databases.GetDB()	

	// querying data
	if err := db.Find(&todos).Error; err != nil {
		log.Printf("Error receiving todos: %v", err)
		return nil, err
	}

	return todos, nil
}

func GetTodosById(id int) (*models.Todo, error) {
	var todoById models.Todo
	db := databases.GetDB()
	if err := db.First(&todoById, id).Error; err != nil {
		log.Printf("Error receiving todos: %v", err)
		return nil, err
	}
	return &todoById, nil
}

func AddTodo(todo *models.Todo) error {
	db := databases.GetDB()

  // check for duplicate title
  var existingTodo models.Todo
  if err := db.Where("title = ?", todo.Title).First(&existingTodo).Error; err != nil{
    return errors.New("Todo already exists")
  }

	result := db.Create(todo)
	return result.Error
}

func UpdateTodo(id int, title string, completed bool) error{
  db := databases.GetDB()
  var todo models.Todo
  db.First(&todo)
  todo.ID = id
  todo.Title = title
  todo.Completed = completed
  result := db.Save(&todo)
  return result.Error
}

func DeleteTodo(id int) error{
  db := databases.GetDB()
  var todo models.Todo 
  todo.ID = id 

  result := db.Delete(&todo)
  return result.Error
}
