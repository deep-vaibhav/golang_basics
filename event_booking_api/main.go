package main

import (
	"example.com/first-app/event_booking_api/db"
	"example.com/first-app/event_booking_api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
