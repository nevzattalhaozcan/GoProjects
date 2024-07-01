package routes

import (
	"log"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		log.Printf("Error fetching events: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events. try again later"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		log.Println("not authorized: empty token")
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	userID, err := utils.VerifyToken(token)

	if err != nil {
		log.Printf("not authorized token: %v", err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	event.UserID = userID

	err = event.Save()
	if err != nil {
		log.Printf("Error saving event: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event. try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Printf("Error parsing event ID: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		log.Printf("Error fetching event by ID: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Printf("Error parsing event ID: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	_, err = models.GetEventByID(eventID)
	if err != nil {
		log.Printf("Error fetching event by ID: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}
	updatedEvent.ID = eventID
	err = updatedEvent.UpdateEvent()
	if err != nil {
		log.Printf("Error updating event: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event. try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event updated", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		log.Printf("Error fetching event by ID: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		log.Printf("Error deleting event: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event. try again later"})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"message": "event deleted", "event": event})
}
