package models



type Appointment struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Title  string `json:"title"`
  Author string `json:"author"`
  Time   string `json:"author"`
}