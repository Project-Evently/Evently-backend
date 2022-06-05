package user

import "github.com/Project-Evently/Evently-backend/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}
func (s *Service) GetUser(uniqueStudentId string) (*entity.User, error) {
	user, err := s.repo.Read(uniqueStudentId)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (s *Service) CreateUser(user *entity.User) (int, error) {
	userId, err := s.repo.Write(user)
	if err != nil {
		return -999, err
	}
	return userId, nil
}

func (s *Service) UpdateUserPassword(uniqueStudentId string, newPassword string) error {
	err := s.repo.UpdatePassword(uniqueStudentId, newPassword)
	if err != nil {
		return err
	}
	return nil
}
