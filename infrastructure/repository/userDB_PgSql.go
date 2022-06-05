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
	rows, err := tx.Query(context.Background(), "SELECT user_id, username, user_password, unique_student_id, college_name, branch_name, current_year, contact_number,email, github_id, event_registered FROM users where unique_student_id = $1", uniqueStudentId)
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
		var CollegeName string
		var BranchName string
		var CurrentYear string
		var ContactNumber string
		var Email string
		var GithubId string

		err = rows.Scan(&UserId, &Username, &Password, &UniqueStudentId, &CollegeName, &BranchName, &CurrentYear, &ContactNumber, &Email, &GithubId, &EventRegistered)
		if err != nil {
			return nil, err
		}
		fmt.Printf("%v\n", Username)
		user = &entity.User{
			UserId:          UserId,
			Username:        Username,
			Password:        Password,
			UniqueStudentId: UniqueStudentId,
			CollegeName:     CollegeName,
			BranchName:      BranchName,
			CurrentYear:     CurrentYear,
			ContactNumber:   ContactNumber,
			Email:           Email,
			GithubId:        GithubId,
			EventRegistered: EventRegistered,
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
	row := tx.QueryRow(context.Background(), "insert into users (username,user_password,unique_student_id, college_name, branch_name, current_year, contact_number,email, github_id) values($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING user_id",
		user.Username, user.Password, user.UniqueStudentId, user.CollegeName, user.BranchName, user.CurrentYear, user.ContactNumber, user.Email, user.GithubId)

	err = row.Scan(&userId)
	if err != nil {
		return -999, err
	}

	tx.Commit(context.Background())

	return userId, nil

}

func (r *UserDbSql) UpdatePassword(uniqueStudentId string, newPassword string) error {
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return err
	}

	ct, err := tx.Exec(context.Background(), "update users set user_password = $1 where user_id = $2", newPassword, uniqueStudentId)
	if err != nil {
		return err
	}
	if ct.RowsAffected() < 1 {
		return errors.New("zero rows affected")
	}

	return nil
}
