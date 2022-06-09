package adminHandlers

import (
	"github.com/Project-Evently/Evently-backend/entity"
	"github.com/Project-Evently/Evently-backend/usecase/admin"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	Service *admin.Service
}

func NewAdminHandler(Service *admin.Service) *AdminHandler {
	return &AdminHandler{
		Service: Service,
	}
}

func (s *AdminHandler) GetInstitute(c *gin.Context) {
	var instituteName = c.Param("instituteName")
	institute, err := s.Service.GetInstituteDetails(instituteName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Unexpected error : ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"institute": institute})
	return
}

func (s *AdminHandler) GetClub(c *gin.Context) {
	var clubName = c.Param("clubName")
	club, err := s.Service.GetClubsDetails(clubName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Unexpected error : ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"club": club})
	return
}

func (s *AdminHandler) CreateInstitute(c *gin.Context) {
	var body entity.Institute
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	institute := &entity.Institute{
		InstituteID:   body.InstituteID,
		InstituteName: body.InstituteName,
		AdminId:       body.AdminId,
		Password:      body.Password,
	}

	err := s.Service.CreateInstituteDetails(institute)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Institute added.": ""})
	return
}

func (s *AdminHandler) CreateClub(c *gin.Context) {
	var body entity.Club
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	club := &entity.Club{
		Id:            body.Id,
		InstituteId:   body.InstituteId,
		ClubName:      body.ClubName,
		ClubPresident: body.ClubPresident,
		AdminId:       body.AdminId,
		AdminPassword: body.AdminPassword,
	}

	err := s.Service.CreateClubsDetails(club)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Club added.": ""})
	return
}

func (s *AdminHandler) GetInstituteList(c *gin.Context) {
	list, err := s.Service.GetInstituteList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"club": list})
	return
}

func (s *AdminHandler) GetClubList(c *gin.Context) {
	var instituteId = c.Param("instituteId")
	id, err := strconv.Atoi(instituteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide a better institute id."})
		return
	}
	list, err := s.Service.GetClubsList(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"club": list})
	return
}
