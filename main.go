package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/Vedantdwivedi/appointment/models" // new

)

func main() {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

  r.Run()
}cd app	