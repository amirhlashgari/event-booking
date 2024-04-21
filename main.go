package main

import (
	"github.com/amirhlashgari/event-booking/db"
	"github.com/amirhlashgari/event-booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
