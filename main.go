package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/Vedantdwivedi/appointment/models" // new
  "github.com/Vedantdwivedi/appointment/controllers"
)

func main() {
  r := gin.Default()

  models.ConnectDatabase() // new
	
  
  r.GET("/list", controllers.FindList) // new
  //r.POST("/books", controllers.CreateBook) // new
  //r.GET("/books/:id", controllers.FindBook) //new
  //r.PATCH("/books/:id", controllers.UpdateBook) // new
  //r.DELETE("/books/:id",controllers.DeleteBook)
  //r.Run()
}