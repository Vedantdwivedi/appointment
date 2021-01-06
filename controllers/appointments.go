package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/Vedantdwivedi/appointment/models"
)

type CreateAppointmentInput struct {
  Name  string `json:"name" binding:"required"`
  Designation string `json:"designation" binding:"required"`
  Time   string `json:"time" binding:"required"`
}

// GET /books
// Get all books
func FindList(c *gin.Context) {
  var appointments []models.Appointment
  models.DB.Find(&appointments)

  c.JSON(http.StatusOK, gin.H{"data": appointments})
}


// POST /Appointment
// Create new appointment
func CreateAppointment(c *gin.Context) {
  // Validate input
  var input CreateAppointmentInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create appointment
  appointment := models.Appointment{Name: input.Name, Designation: input.Designation, Time: input.Time}
  models.DB.Create(&appointment)

  c.JSON(http.StatusOK, gin.H{"data": appointment})
}




// POST /Schedule
// Create doctor schedule
func CreateSchedule(c *gin.Context) {
  // Validate input
  var input CreateAppointmentInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create appointment
  appointment := models.Appointment{Name: input.Name, Designation: input.Designation, Time: input.Time}
  models.DB.Create(&appointment)

  c.JSON(http.StatusOK, gin.H{"data": appointment})
}

// GET /doctor appointment/:designation
// Find a book
func FindLists(c *gin.Context) {  // Get model if exist
  var appointment models.Appointment

  if err := models.DB.Where("name = ?", c.Param("name")).Find(&appointment).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	return
   	
  }

  c.JSON(http.StatusOK, gin.H{"data": appointment})
}



// CANCEL /Appointment/:id
// Cancel a appointment
func CancelAppointment(c *gin.Context) {
  // Get model if exist
  var appointment models.Appointment
  if err := models.DB.Where("id = ?", c.Param("id")).First(&appointment).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&appointment)

  c.JSON(http.StatusOK, gin.H{"data": true})
}