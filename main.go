package main

import (
	"database/sql"
	"finalproject-sanber-soni/database"
	"fmt"
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

	StartServer()
}
