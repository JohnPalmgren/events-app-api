package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

// registerForEvent handles the registration of a user for a specified event.
// It expects the event ID as a URL parameter and the user ID to be set in the context.
// If the event ID is invalid, the event cannot be fetched, or the registration fails,
// it returns an appropriate HTTP status code and error message.
func registerForEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Data - eventId"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	// Set during authentication
	userId := context.GetInt64("userId")

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

// cancelRegistration handles the cancellation of a user's registration for a specified event.
// It expects the event ID as a URL parameter and the user ID to be set in the context.
// If the event ID is invalid, the event cannot be found, or the cancellation fails,
// it returns an appropriate HTTP status code and error message.
func cancelRegistration(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Data - eventId"})
		return
	}

	var event models.Event
	event.ID = eventId

	// Set during authentication
	userId := context.GetInt64("userId")

	err = event.DeleteRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel user registration for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled"})
}
