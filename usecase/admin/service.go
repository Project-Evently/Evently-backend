package admin

import "github.com/Project-Evently/Evently-backend/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetInstituteDetails(instituteName string) (*entity.Institute, error) {
	institute, err := s.repo.ReadInstitute(instituteName)
	if err != nil {
		return nil, err
	}
	return institute, nil
}

func (s *Service) GetClubsDetails(clubName string) (*entity.Club, error) {
	club, err := s.repo.ReadClubs(clubName)
	if err != nil {
		return nil, err
	}
	return club, nil
}

func (s *Service) CreateInstituteDetails(institute *entity.Institute) error {
	err := s.repo.WriteInstitute(institute)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CreateClubsDetails(club *entity.Club) error {
	err := s.repo.WriteClubs(club)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetInstituteList() ([]*entity.Institute, error) {
	list, err := s.repo.ReadInstituteList()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Service) GetClubsList(instituteId int) ([]*entity.Club, error) {
	list, err := s.repo.ReadClubsList(instituteId)
	if err != nil {
		return nil, err
	}
	return list, nil
}
