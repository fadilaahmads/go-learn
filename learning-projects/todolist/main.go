package main

import (
	"gotodolist/databases"
	"gotodolist/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
  // load env variables from .env
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Failed to load .env: ", err)
  }

  // initialize database
  databases.Init()

  // initialize gin router
  r := gin.Default()

  // setup routes
  routes.SetupRoutes(r)
  
  r.Run(":8080")
}
