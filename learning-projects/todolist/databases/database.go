package databases

import (
  "log"
  "fmt"
  "os"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
/* 	"gotodolist/models" */
)

var DB *gorm.DB 

func Init() *gorm.DB {
  dbUser := os.Getenv("DB_USER")
  dbPassword := os.Getenv("DB_PASS")
  dbHost := os.Getenv("DB_HOST")
  dbName := os.Getenv("DB_NAME")
  dbPort := os.Getenv("DB_PORT")

  // create a connection string
  dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

  // open connection to database 
  DB, err := gorm.Open("mysql", dbConnectionString)
  if err != nil {
    log.Fatal("Failed to connect to database: ", err)
  }

  // Automatically migrate the schema, to keep your schema updated.
  //DB.AutoMigrate(&models.Todo{})
  return DB
}

func GetDB() *gorm.DB{
  db := Init()
  if db == nil {
    log.Panicln("Failed to initialize database connection")
  }
  return db
 }
