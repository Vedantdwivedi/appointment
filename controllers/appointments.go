package controllers

import (
  "github.com/gin-gonic/gin"
  "github.com/Vedantdwivedi/appointment/models"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
  var appointments []models.Appointment
  models.DB.Find(&appointments)

  c.JSON(http.StatusOK, gin.H{"data": appointments})
}