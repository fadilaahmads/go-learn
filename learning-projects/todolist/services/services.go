package services

import (
	"fmt"
	"gotodolist/databases"
	"gotodolist/models"
	"log"

	/*   "github.com/jinzhu/gorm" */
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/* var db *gorm.DB = databases.GetDB()   */

func GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	db := databases.GetDB()
	if db == nil {
		log.Println("Database connection is not initialize")
		return nil, fmt.Errorf("Database is not initialize")
	}

	// querying data
	if err := db.Find(&todos).Error; err != nil {
		log.Printf("Error receiving todos: %v", err)
		return nil, err
	}

	return todos, nil
}

func GetTodosById(id int) ([]models.Todo, error) {
	var todoById []models.Todo
	db := databases.GetDB()
	if err := db.First(&todoById, id).Error; err != nil {
		log.Printf("Error receiving todos: %v", err)
		return nil, err
	}
	return todoById, nil
}

func AddTodo(todo models.Todo) error {
	db := databases.GetDB()
	result := db.Create(todo)
	return result.Error
}
