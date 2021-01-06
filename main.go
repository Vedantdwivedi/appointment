package main

import (
  "github.com/gin-gonic/gin"
  "github.com/Vedantdwivedi/appointment/models" 
  "github.com/Vedantdwivedi/appointment/controllers"
    //"github.com/rahmanfadhil/gin-bookstore/controllers"
	//"github.com/rahmanfadhil/gin-bookstore/models"
)

func main() {
  r := gin.Default()

  models.ConnectDataBase() // new
	
  
  r.GET("/list", controllers.FindList) // new
  //r.POST("/books", controllers.CreateBook) // new
  //r.GET("/books/:id", controllers.FindBook) //new
  //r.PATCH("/books/:id", controllers.UpdateBook) // new
  //r.DELETE("/books/:id",controllers.DeleteBook)
  r.Run()
}