package main

import (
	"database/sql"
	"finalproject-sanber-soni/auth"
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

	r := gin.Default()

	// user endpoint handler
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := controllers.NewUserHandler(userService)
	users := r.Group("/users")
	// all user
	users.POST("/register", userHandler.RegisterUser)
	users.POST("/login", userHandler.Login)
	users.PUT("/edit", auth.MiddlewareUserAuth(userService), userHandler.UpdateUser)
	users.DELETE("/delete", auth.MiddlewareUserAuth(userService), userHandler.DeleteUser)
	// admin
	users.GET("/get-all-users", auth.MiddlewareUserAuth(userService), userHandler.GetAllUsers)

	// categories endpoint handler
	categoryRepository := repository.NewCategoriesRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := controllers.NewCatHandler(categoryService)
	category := r.Group("/category")
	// all user
	category.GET("/get-all", categoryHandler.GetAllCategories)
	category.GET("/:category_id/items", categoryHandler.GetAllInventoriesByCatId)
	// admin
	category.POST("/add", auth.MiddlewareUserAuth(userService), categoryHandler.InsertCategory)
	category.PUT("/edit/:category_id", auth.MiddlewareUserAuth(userService), categoryHandler.UpdateCategory)
	category.DELETE("/delete/:category_id", auth.MiddlewareUserAuth(userService), categoryHandler.DeleteCategories)

	// invenotry endpoint handler
	inventoryRepository := repository.NewInventoryRepository(db)
	inventoryService := service.NewInventoryService(inventoryRepository)
	inventoryHandler := controllers.NewInventoryHandler(inventoryService)
	inventory := r.Group("/inventory")
	// all user
	inventory.GET("/get-all", inventoryHandler.GetAll)
	inventory.GET("/get/:inven_id", inventoryHandler.GetById)
	// admin
	inventory.POST("/add", auth.MiddlewareUserAuth(userService), inventoryHandler.InsertInventory)
	inventory.PUT("/edit/:inven_id", auth.MiddlewareUserAuth(userService), inventoryHandler.UpdateInventory)
	inventory.DELETE("/delete/:inven_id", auth.MiddlewareUserAuth(userService), inventoryHandler.DeleteInventory)

	// transactions endpoint handler
	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)
	transactionHandler := controllers.NewTransactionHandler(transactionService)
	transaction := users.Group("/transaction")
	// user
	transaction.GET("/get-all", auth.MiddlewareUserAuth(userService), transactionHandler.GetAll)
	// /get?status=Paid
	transaction.GET("/get", auth.MiddlewareUserAuth(userService), transactionHandler.GetByStatus)
	transaction.POST("/create", auth.MiddlewareUserAuth(userService), transactionHandler.CreateTransaction)
	// misal /1?action=pay atau /1?action=cancel
	transaction.PUT("/:trans_id", auth.MiddlewareUserAuth(userService), transactionHandler.UpdateTransaction)
	transaction.PUT("/admin/:trans_id", auth.MiddlewareUserAuth(userService), transactionHandler.UpdateAdmin)
	r.Run("127.0.0.1:8080")
}
