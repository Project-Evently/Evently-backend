package api

import (
	"context"
	"fmt"
	"github.com/Project-Evently/Evently-backend/api/handler/adminHandlers"
	"github.com/Project-Evently/Evently-backend/api/handler/userHandlers"
	"github.com/Project-Evently/Evently-backend/infrastructure/repository"
	"github.com/Project-Evently/Evently-backend/usecase/admin"
	"github.com/Project-Evently/Evently-backend/usecase/user"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

func Server() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	/*dbUrl := os.Getenv("DATABASE_URL")

	if dbUrl == "" {
		return
	}*/
	REMOTE_DB := "postgres://lenpqawjszuuyo:6c22194378e865810e22e1abb7fa79ec60f89d91b0c0f454a853aeb3da7e7269@ec2-52-30-67-143.eu-west-1.compute.amazonaws.com:5432/dab194a6cpbrq1"
	//	KEY_DATABASE_URL := "postgres://Vishwajeet:vishvapriya123@localhost:5432/EventDb?&pool_max_conns=10"
	//	DATABASE_URL := os.Getenv("DATABASE_URL")
	//	DATABASE_URL := KEY_DATABASE_URL
	fmt.Printf(" DB = %s\n", REMOTE_DB)
	ctx := context.Background()
	cofig, err := pgxpool.ParseConfig(REMOTE_DB)
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

	AdminRepo := repository.NewAdminDbSql(pool)
	AdminService := admin.NewService(AdminRepo)
	AdminHandler := adminHandlers.NewAdminHandler(AdminService)

	address := fmt.Sprintf("0.0.0.0:%v", port)
	fmt.Printf("%v", address)
	r := gin.Default()
	r.POST("/api/v1/user/create", UserHandler.CreateUser)
	r.GET("/api/v1/user/:uniqueStudentId", UserHandler.GetUser)
	r.POST("api/v1/user/updatePassword", UserHandler.UpdatePassword)
	r.GET("/api/v1/institute/:instituteName", AdminHandler.GetInstitute)
	r.GET("/api/v1/institute/list", AdminHandler.GetInstituteList)
	r.GET("/api/v1/club/:clubName", AdminHandler.GetClub)
	r.GET("/api/v1/club/list", AdminHandler.GetClubList)
	r.POST("/api/v1/institute/create", AdminHandler.CreateInstitute)
	r.POST("/api/v1/club/create", AdminHandler.CreateClub)
	r.Run(address)
}
