package models



type Appointment struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Name  string `json:"name"`
  Designation string `json:"designation"`
  Time   string `json:"time"`
}