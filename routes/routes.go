package routes

import (
	"example.com/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)

	server.POST("/login", login)

	server.GET("/events", getEvents)

	server.GET("event/:id", getEvent)
	server.POST("/event", middleware.Authenticate, createEvents)
	server.POST("/event/:id/register", middleware.Authenticate, registerForEvent)
	server.PUT("/event/:id", middleware.Authenticate, updateEvent)
	server.DELETE("/event/:id", middleware.Authenticate, deleteEvent)
	server.DELETE("/event/:id/register", middleware.Authenticate, cancelRegistration)
}
