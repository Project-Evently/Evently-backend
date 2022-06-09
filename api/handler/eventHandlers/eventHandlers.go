package eventHandlers

import (
	"github.com/Project-Evently/Evently-backend/entity"
	"github.com/Project-Evently/Evently-backend/usecase/event"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EventHandler struct {
	Service *event.Service
}

func NewEventHandler(Service *event.Service) *EventHandler {
	return &EventHandler{
		Service: Service,
	}
}

func (s *EventHandler) GetEventById(c *gin.Context) {
	var eventId = c.Param("eventId")
	id, err := strconv.Atoi(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Unexpected error : ": err.Error()})
		return
	}
	event, err := s.Service.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Unexpected error : ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
	return
}

func (s *EventHandler) GetEventListByInstitute(c *gin.Context) {
	var eventId = c.Param("instituteId")
	id, err := strconv.Atoi(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Unexpected error : ": err.Error()})
		return
	}
	events, err := s.Service.GetEventListByInstitute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Unexpected error : ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": events})
	return
}

func (s *EventHandler) GetEventListByClub(c *gin.Context) {
	var clubId = c.Param("clubId")
	id, err := strconv.Atoi(clubId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Unexpected error : ": err.Error()})
		return
	}
	events, err := s.Service.GetEventListByClub(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Unexpected error : ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": events})
	return
}

func (s *EventHandler) CreateEvent(c *gin.Context) {
	var body entity.Event
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event := &entity.Event{
		EventId:        body.EventId,
		ClubId:         body.ClubId,
		InstituteId:    body.InstituteId,
		Description:    body.Description,
		EventDateIST:   body.EventDateIST,
		EventTimeIST:   body.EventTimeIST,
		EventLocation:  body.EventLocation,
		EventOrganizer: body.EventOrganizer,
		EventContact:   body.EventContact,
		EventLink:      body.EventLink,
	}

	err := s.Service.CreateEvent(event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Event added.": ""})
	return
}
