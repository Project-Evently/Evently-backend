package event

import "github.com/Project-Evently/Evently-backend/entity"

type Reader interface {
	ReadEvent(eventId int) (*entity.Event, error)
	ReadListInstitute(instituteId int) ([]*entity.Event, error)
	ReadListClub(clubId int) ([]*entity.Event, error)
}

type Writer interface {
	WriteEvent(event *entity.Event) error
}

type Repository interface {
	Reader
	Writer
}

type Usecase interface {
	GetEventById(eventId int) (*entity.Event, error)
	GetEventListByInstitute(instituteId int) ([]*entity.Event, error)
	GetEventListByClub(clubId int) ([]*entity.Event, error)
	CreateEvent(event *entity.Event) error
}
