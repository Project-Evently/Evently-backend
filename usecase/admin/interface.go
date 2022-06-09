package admin

import "github.com/Project-Evently/Evently-backend/entity"

type Reader interface {
	ReadInstitute(instituteName string) (*entity.Institute, error)
	ReadClubs(clubName string) (*entity.Club, error)
	ReadInstituteList() ([]*entity.Institute, error)
	ReadClubsList(instituteId int) ([]*entity.Club, error)
}

type Writer interface {
	WriteInstitute(institute *entity.Institute) error
	WriteClubs(club *entity.Club) error
}

type Repository interface {
	Reader
	Writer
}

type Usecase interface {
	GetInstituteDetails(instituteName string) (*entity.Institute, error)
	GetClubsDetails(clubName string) (*entity.Club, error)
	CreateInstituteDetails(institute *entity.Institute) error
	CreateClubsDetails(club *entity.Club) error
	GetInstituteList() ([]*entity.Institute, error)
	GetClubsList(instituteId int) ([]*entity.Club, error)
}
