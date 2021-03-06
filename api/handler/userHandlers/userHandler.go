package userHandlers

import (
	"github.com/Project-Evently/Evently-backend/entity"
	"github.com/Project-Evently/Evently-backend/usecase/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	Service *user.Service
}

func NewUserHandler(Service *user.Service) *UserHandler {
	return &UserHandler{
		Service: Service,
	}
}

func (s *UserHandler) GetUser(c *gin.Context) {
	var UniqueStudentId = c.Param("uniqueStudentId")
	user, err := s.Service.GetUser(UniqueStudentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Unexpected error : ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (s *UserHandler) CreateUser(c *gin.Context) {
	var body entity.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := &entity.User{
		UserId:          body.UserId,
		Username:        body.Username,
		Password:        body.Password,
		UniqueStudentId: body.UniqueStudentId,
		CollegeName:     body.CollegeName,
		BranchName:      body.BranchName,
		CurrentYear:     body.CurrentYear,
		ContactNumber:   body.ContactNumber,
		Email:           body.Email,
		GithubId:        body.GithubId,
		EventRegistered: body.EventRegistered,
	}

	userId, err := s.Service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"UserId": userId})
	return
}

func (s *UserHandler) UpdatePassword(c *gin.Context) {
	var body entity.UpdatePassword
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//var UniqueStudentId = c.Param("uniqueStudentId")
	err := s.Service.UpdateUserPassword(body.UniqueStudentId, body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Unexpected error : ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Response": "Password updated"})
}
