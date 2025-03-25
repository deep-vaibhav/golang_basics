package main

import (
	"fmt"
	"net/http"

	"example.com/first-app/event_booking_api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

// var DB *sql.DB

// func InitDB() {
//     var err error
//     DB, err = sql.Open("sqlite3", "api.db")

//     if err != nil {
//         panic("Could not connect to database.")
//     }

//     DB.SetMaxOpenConns(10)
//     DB.SetMaxIdleConns(5)

//     createTables()
// }
