package user

import "github.com/Project-Evently/Evently-backend/entity"

type Reader interface {
	Read(uniqueStudentId string) (*entity.User, error)
}

type Writer interface {
	Write(user *entity.User) (int, error)
}

type Updater interface {
	//UpdateSocialLinks(uniqueStudentId string, links []string) error
	UpdatePassword(uniqueStudentId string, newPassword string) error
	//UpdateEventRegistered(uniqueStudentId string, events []string) error
	//UpdateUsername(uniqueStudentId string, username string) error
}

/*type Deleter interface {
	Delete(userId int) error
}*/
type Repository interface {
	Reader
	Writer
	Updater
	//Deleter
}
type Usecase interface {
	GetUser(uniqueStudentId string) (*entity.User, error)
	CreateUser(user *entity.User) (int, error)
	UpdateUserPassword(uniqueStudentId string, newPassword string) error
	//UpdateUserUsername(uniqueStudentId string, username string) error
	//UpdateUserEventRegistered(uniqueStudentId string, events []string) error
	//UpdateUserSocialLinks(uniqueStudentId string, links []string) error
	//	DeleteUser(userId int) error
}
