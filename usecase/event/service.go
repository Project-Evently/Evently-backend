package event

import "github.com/Project-Evently/Evently-backend/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetEventById(eventId int) (*entity.Event, error) {
	event, err := s.repo.ReadEvent(eventId)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (s *Service) GetEventListByInstitute(instituteId int) ([]*entity.Event, error) {
	events, err := s.repo.ReadListInstitute(instituteId)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *Service) GetEventListByClub(clubId int) ([]*entity.Event, error) {
	events, err := s.repo.ReadListClub(clubId)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *Service) CreateEvent(event *entity.Event) error {
	err := s.repo.WriteEvent(event)
	if err != nil {
		return err
	}
	return nil
}
