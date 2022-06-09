package repository

import (
	"context"
	"errors"
	"github.com/Project-Evently/Evently-backend/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AdminDbSql struct {
	pool *pgxpool.Pool
}

func NewAdminDbSql(pool *pgxpool.Pool) *AdminDbSql {
	return &AdminDbSql{
		pool: pool,
	}
}

func (r *AdminDbSql) ReadInstitute(instituteFullName string) (*entity.Institute, error) {
	var Institute *entity.Institute
	//Read from Db
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	row := tx.QueryRow(context.Background(), "select institute_id, institute_name, admin_id, admin_password from institutions where institute_name = $1", instituteFullName)
	var instituteID string
	var instituteName string
	var adminId string
	var password string

	err = row.Scan(&instituteID, &instituteName, &adminId, &password)
	if err != nil {
		return nil, err
	}

	Institute = &entity.Institute{
		InstituteID:   instituteID,
		InstituteName: instituteFullName,
		AdminId:       adminId,
		Password:      password,
	}

	if Institute == nil {
		return nil, errors.New("unable to fetch Institute details")
	}
	tx.Commit(context.Background())

	return Institute, nil
}

func (r *AdminDbSql) WriteInstitute(institute *entity.Institute) error {
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return err
	}

	ct, err := tx.Exec(context.Background(), "insert into institutions (institute_name,admin_id,admin_password) values ($1,$2,$3)", institute.InstituteName, institute.AdminId, institute.Password)
	if err != nil || ct.RowsAffected() < 0 {
		return errors.New("unable to Write data to institutions table")
	}

	tx.Commit(context.Background())
	return nil
}

func (r *AdminDbSql) ReadClubs(clubName string) (*entity.Club, error) {
	var club *entity.Club

	var Id int
	var InstituteId int
	var ClubName string
	var ClubPresident string
	var AdminId string
	var AdminPassword string
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	row := tx.QueryRow(context.Background(), "select club_id, institute_id, club_name, club_president,admin_id,admin_password from clubs where club_name = $1", clubName)
	err = row.Scan(&Id, &InstituteId, &ClubName, &ClubPresident, &AdminId, &AdminPassword)
	club = &entity.Club{
		Id:            Id,
		InstituteId:   InstituteId,
		ClubName:      ClubName,
		ClubPresident: ClubPresident,
		AdminId:       AdminId,
		AdminPassword: AdminPassword,
	}
	if err != nil || club == nil {
		return nil, err
	}

	return club, nil
}

func (r *AdminDbSql) WriteClubs(club *entity.Club) error {
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return err
	}
	ct, err := tx.Exec(context.Background(), "insert into clubs (institute_id, club_name, club_president,admin_id,admin_password) values ($1,$2,$3,$4,$5)", club.InstituteId, club.ClubName, club.ClubPresident, club.AdminId, club.AdminPassword)
	if err != nil || ct.RowsAffected() < 0 {
		return errors.New("unable to Write data to Clubs table")
	}
	tx.Commit(context.Background())
	return nil
}

func (r *AdminDbSql) ReadInstituteList() ([]*entity.Institute, error) {

	InstituteList := make([]*entity.Institute, 1)

	var Institute *entity.Institute
	var instituteID string
	var instituteName string

	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query(context.Background(), "select institute_id, institute_name from institutions")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&instituteID, &instituteName)
		if err != nil {
			return nil, err
		}
		Institute = &entity.Institute{
			InstituteID:   instituteID,
			InstituteName: instituteName,
		}
		InstituteList = append(InstituteList, Institute)
	}

	return InstituteList, nil
}

func (r *AdminDbSql) ReadClubsList(instituteId int) ([]*entity.Club, error) {
	ClubsList := make([]*entity.Club, 1)

	var club *entity.Club

	var Id int
	var InstituteId int
	var ClubName string

	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(context.Background(), "select club_id, institute_id, club_name from clubs where institute_id = $1 ", instituteId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&Id, &instituteId, &ClubName)
		if err != nil {
			return nil, err
		}
		club = &entity.Club{
			Id:          Id,
			InstituteId: InstituteId,
			ClubName:    ClubName,
		}
		ClubsList = append(ClubsList, club)
	}

	return ClubsList, nil

}
