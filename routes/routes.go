package routes

import (
	"github.com/amirhlashgari/event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getEvent)
	// server.POST("/events", middlewares.Authenticate, createEvent) ---> one way of create protected routes

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:eventId", updateEvent)
	authenticated.DELETE("/events/:eventId", deleteEvent)

	authenticated.POST("/events/:eventId/register", registerForEvent)
	authenticated.DELETE("/events/:eventId/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
