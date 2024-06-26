package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

// getEvents retrieves all events from the database.
// If there is an error fetching the events, it returns an internal server error status.
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

// getEvent retrieves a specific event by its ID from the database.
// It expects the event ID as a URL parameter. If the event ID is invalid or there is an error
// fetching the event, it returns an internal server error status.
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Data - event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	context.JSON(http.StatusOK, event)

}

// createEvents handles the creation of a new event.
// It expects the event data in the request body and the user ID to be set in the context.
// If the input data is invalid or there is an error creating the event, it returns an appropriate
// HTTP status code and error message.
func createEvents(context *gin.Context) {

	// Set during authentication
	userId := context.GetInt64("userId")

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Data - event"})
		return
	}

	event.UserId = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

// updateEvent handles updating an existing event.
// It expects the event ID as a URL parameter, the updated event data in the request body,
// and the user ID to be set in the context. If the event ID is invalid, the input data is invalid,
// or there is an error updating the event, it returns an appropriate HTTP status code and error message.
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Data - eventId"})
		return
	}

	// Set during authentication
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Data - event"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

// deleteEvent handles the deletion of an existing event.
// It expects the event ID as a URL parameter and the user ID to be set in the context.
// If the event ID is invalid or there is an error deleting the event,
// it returns an appropriate HTTP status code and error message.
func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Data - event ID"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Event does not exist"})
		return
	}

	if userId != event.UserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
