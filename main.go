package main

import (
	"database/sql"
	"finalproject-sanber-soni/controllers"
	"finalproject-sanber-soni/database"
	"finalproject-sanber-soni/repository"
	"finalproject-sanber-soni/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

var (
	db  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		log.Fatal(err)
	}
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		"localhost",
		5432,
		"postgres",
		"postgres",
		"cobafp2")

	db, err = sql.Open("postgres", psqlInfo)
	err = db.Ping()
	if err != nil {
		fmt.Println("Connection to database is failed")
		panic(err)
	}

	fmt.Println("Successfully make connection to database")

	database.DbMigrate(db)
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := controllers.NewUserHandler(userService)

	r := gin.Default()

	r.POST("/users/register", userHandler.RegisterUser)
	r.POST("/users/login", userHandler.Login)
	r.PUT("/users/edit/:user_id", userHandler.UpdateUser)

	r.Run("127.0.0.1:8080")
}
