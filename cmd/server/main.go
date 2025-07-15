package main

import (
	"log"
	nativeHttp "net/http"

	redisLib "golang_test/internal/shared/redis"

	authHttp "golang_test/internal/auth/delivery/http"
	authRepo "golang_test/internal/auth/repository"
	authUsecase "golang_test/internal/auth/usecase"

	userHttp "golang_test/internal/user/delivery/http"
	userRepo "golang_test/internal/user/repository"
	userUsecase "golang_test/internal/user/usecase"

	"golang_test/internal/shared/config"
	"golang_test/internal/shared/jwt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadDBConfig()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	db, err := sqlx.Connect("postgres", cfg.DSN())
	if err != nil {
		log.Fatal("failed to connect db:", err)
	}

	redisLib.InitRedis("localhost:6379", "", 0)

	// auth
	authRepository := authRepo.NewAuthRepository(db)
	authUc := authUsecase.NewAuthUsecase(authRepository)
	authHandler := authHttp.NewAuthHandler(authUc)

	// user
	userRepository := userRepo.NewUserRepository(db)
	userUc := userUsecase.NewUserUsecase(userRepository)
	userHandler := userHttp.NewUserHandler(userUc)

	router := gin.Default()

	router.POST("/auth/login", authHandler.Login)

	userGroup := router.Group("/users")
	userGroup.Use(jwt.AuthMiddleware())
	{
		userGroup.GET("/user", userHandler.GetAll)
		userGroup.POST("/user", userHandler.Create)
		userGroup.PUT("/user", userHandler.Update)
		userGroup.DELETE("/user/:id", userHandler.Delete)
	}

	log.Println("server running at :8080")
	if err := nativeHttp.ListenAndServe(":8080", router); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
