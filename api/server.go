package api

import (
	"context"
	"fmt"
	"github.com/Project-Evently/Evently-backend/api/handler/userHandlers"
	"github.com/Project-Evently/Evently-backend/infrastructure/repository"
	"github.com/Project-Evently/Evently-backend/usecase/user"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Server() {

	KEY_DATABASE_URL := "postgres://Vishwajeet:vishvapriya123@localhost:5432/EventDb?&pool_max_conns=10"
	//	DATABASE_URL := os.Getenv("DATABASE_URL")
	DATABASE_URL := KEY_DATABASE_URL
	fmt.Printf(" DB = %s\n", DATABASE_URL)
	ctx := context.Background()
	cofig, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, cofig)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}
	defer pool.Close()

	UserRepo := repository.NewUserDbSql(pool)
	UserService := user.NewService(UserRepo)
	UserHandler := userHandlers.NewUserHandler(UserService)

	r := gin.Default()
	r.POST("/api/v1/user/create", UserHandler.CreateUser)
	r.GET("/api/v1/user/:uniqueStudentId", UserHandler.GetUser)
	r.Run(":8080")
}
