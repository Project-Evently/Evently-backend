package user

import "github.com/Project-Evently/Evently-backend/entity"

type Reader interface {
	Read(uniqueStudentId string) (*entity.User, error)
}

type Writer interface {
	Write(user *entity.User) (int, error)
}

/*type Deleter interface {
	Delete(userId int) error
}*/
type Repository interface {
	Reader
	Writer
	//Deleter
}
type Usecase interface {
	GetUser(uniqueStudentId string) (*entity.User, error)
	CreateUser(user *entity.User) (int, error)
	//	DeleteUser(userId int) error
}
