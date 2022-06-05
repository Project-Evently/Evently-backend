package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Project-Evently/Evently-backend/entity"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type UserDbSql struct {
	pool *pgxpool.Pool
}

func NewUserDbSql(pool *pgxpool.Pool) *UserDbSql {
	return &UserDbSql{
		pool: pool,
	}
}

func (r *UserDbSql) Read(uniqueStudentId string) (*entity.User, error) {
	var user *entity.User
	//Read from Db
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query(context.Background(), "SELECT user_id, username, user_password, unique_student_id, event_registered, social_links FROM users where unique_student_id = $1", uniqueStudentId)
	if err != nil {
		log.Printf("UserDB : Error while fetching user details \nError : %s ", err.Error())
		return nil, err
	}

	for rows.Next() {
		var UserId int
		var Username string
		var Password string
		var UniqueStudentId string
		var EventRegistered []int
		var Social []string

		err = rows.Scan(&UserId, &Username, &Password, &UniqueStudentId, &EventRegistered, &Social)
		if err != nil {
			return nil, err
		}
		fmt.Printf("%v\n", Username)
		user = &entity.User{
			UserId:          UserId,
			Username:        Username,
			Password:        Password,
			UniqueStudentId: UniqueStudentId,
			EventRegistered: EventRegistered,
			Social:          Social,
		}
	}

	if user == nil {
		return nil, errors.New("error while fetching data")
	}
	rows.Close()
	tx.Commit(context.Background())

	return user, nil

}

func (r *UserDbSql) Write(user *entity.User) (int, error) {
	userId := -999
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return -999, err
	}
	row := tx.QueryRow(context.Background(), "insert into users (username,user_password,unique_student_id,social_links) values($1,$2,$3,$4) RETURNING user_id", user.Username, user.Password, user.UniqueStudentId, user.Social)

	err = row.Scan(&userId)
	if err != nil {
		return -999, err
	}

	tx.Commit(context.Background())

	return userId, nil

}
