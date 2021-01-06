# ResT API with Go Lang , GORM, Gin Gonic, SQL Lite

This Api performs following tasks:

GET("/list")  - Give List of all the Doctors and Patients booked appointments
GET("/lists/:name") - Give List of Doctor or Patient by name 
POST("/book") - Enable Patuent to book apointment 
POST("/schedule") -Enable doctor to specify what time he/she is available for appointments
DELETE("/cancel/:id") - Enable Doctor/Patient to cancel the appointment 

# Body in JSON for POST

    {
        "name": "Mukesh",
        "designation": "Doctor",
        "time": "09:00 AM - 2:00 PM"
    }
    
##Note: for "name" & "designation" use the notation "Ram"/ "Doctor" (initials should be capital) it shall be easy to query later.

##For further information please do visit the demo: https://youtu.be/kLmThX7ulCs

# Area Of Improvement:
1.) Need to apply authentication & authorisation
2.) Implement date time module for handling appointment
3.) Proper User management 
