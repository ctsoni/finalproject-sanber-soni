package main

import (
	"finalproject-sanber-soni/auth"
	"finalproject-sanber-soni/controllers"
	"finalproject-sanber-soni/repository"
	"finalproject-sanber-soni/service"
	"github.com/gin-gonic/gin"
)

func StartServer() {

	r := gin.Default()
	users := r.Group("/users")
	category := r.Group("/category")
	inventory := r.Group("/inventory")
	transaction := users.Group("/transaction")
	review := r.Group("/review")

	// user endpoint handler
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := controllers.NewUserHandler(userService)
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
	// user
	transaction.GET("/get-all", auth.MiddlewareUserAuth(userService), transactionHandler.GetAll)
	// example /get?status=Paid
	transaction.GET("/get", auth.MiddlewareUserAuth(userService), transactionHandler.GetByStatus)
	transaction.POST("/create", auth.MiddlewareUserAuth(userService), transactionHandler.CreateTransaction)
	// example /1?action=pay or /1?action=cancel
	transaction.PUT("/:trans_id", auth.MiddlewareUserAuth(userService), transactionHandler.UpdateTransaction)
	transaction.PUT("/admin/:trans_id", auth.MiddlewareUserAuth(userService), transactionHandler.UpdateAdmin)
	// admin
	transaction.GET("/admin/get-all", auth.MiddlewareUserAuth(userService), transactionHandler.GetAllAdmin)
	transaction.GET("/admin/get", auth.MiddlewareUserAuth(userService), transactionHandler.GetByStatusAdmin)

	// reviews endpoint handler
	reviewRepository := repository.NewReviewRepository(db)
	reviewService := service.NewReviewService(reviewRepository)
	reviewHanlder := controllers.NewReviewHandler(reviewService)
	// user
	review.GET("/get-all", reviewHanlder.GetAll)
	review.GET("/get-by-user/:user_id", reviewHanlder.GetByUserID)
	review.GET("/get-by-inven/:inven_id", reviewHanlder.GetByInvenID)
	review.POST("/add/:trans_id", auth.MiddlewareUserAuth(userService), reviewHanlder.AddReview)
	review.PUT("/edit/:review_id", auth.MiddlewareUserAuth(userService), reviewHanlder.EditReview)
	review.DELETE("/delete/:review_id", auth.MiddlewareUserAuth(userService), reviewHanlder.DeleteReview)

	r.Run("127.0.0.1:8080")
}
