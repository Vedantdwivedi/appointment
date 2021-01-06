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
  r.GET("/lists/:name", controllers.FindLists) // new 
  r.POST("/book", controllers.CreateAppointment) // new
  r.POST("/schedule", controllers.CreateSchedule) // new
  r.DELETE("/cancel/:id",controllers.CancelAppointment)
  r.Run()
}