package repository

import "github.com/jackc/pgx/v4/pgxpool"

type InstitutionDbSql struct {
	pool *pgxpool.Pool
}

func NewInstitutionDbSql(pool *pgxpool.Pool) *UserDbSql {
	return &UserDbSql{
		pool: pool,
	}
}
