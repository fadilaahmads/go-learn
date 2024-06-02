package utils

import "github.com/gin-gonic/gin"

func ResponseWithJSON(c *gin.Context, code int, payload interface{}){
  c.JSON(code, payload)
}

func ResponseWithError(c *gin.Context, code int, message string){
  ResponseWithJSON(c, code, map[string]string{"error":message})
}
