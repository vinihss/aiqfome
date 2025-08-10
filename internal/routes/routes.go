package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vinihss/aiqfome/internal/infrastructure/database/repositories"
	http_interfaces_authentication "github.com/vinihss/aiqfome/internal/interfaces/http/authentcation"
	customeruse "github.com/vinihss/aiqfome/internal/usecases/customer"
	favoriteuse "github.com/vinihss/aiqfome/internal/usecases/favorite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/vinihss/aiqfome/docs"
	"github.com/vinihss/aiqfome/internal/interfaces/http/customer"
	"github.com/vinihss/aiqfome/internal/interfaces/http/favorite"
	"github.com/vinihss/aiqfome/middlewares"
)

// SetupRoutes @title Aiqfome API
func SetupRoutes(router *gin.Engine) {

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authController := http_interfaces_authentication.NewAuthenticationController()
	router.POST("/authenticate", http_interfaces_authentication.NewAuthenticationHandler(authController).Authenticate)
	db, _ := gorm.Open(postgres.Open("host=0.0.0.0 user=postgres password=postgres dbname=favorites port=5432 sslmode=disable"), &gorm.Config{})
	authorized := router.Group("/")
	authorized.Use(middlewares.JWTAuth())
	{

		favRepo := repositories.NewFavoriteRepository(db)
		createFavoriteUC := favoriteuse.NewCreateFavoriteUseCase(favRepo)
		favController := http_interfaces_favorite.NewFavoriteController(createFavoriteUC)
		favHandler := http_interfaces_favorite.NewFavoriteHandler(favController)
		authorized.POST("/favorites", favHandler.AddFavorite)

		custRepo := repositories.NewCustomerRepository(db)
		createCustomerUC := customeruse.NewCreateCustomerUseCase(custRepo)
		deleteCustomerUC := customeruse.NewDeleteCustomerUseCase(custRepo)
		findCustomerUC := customeruse.NewFindCustomerUseCase(custRepo)
		updateCustomerUC := customeruse.NewUpdateCustomerUseCase(custRepo)

		custController := http_interfaces_customer.NewCustomerController(createCustomerUC, deleteCustomerUC, findCustomerUC, updateCustomerUC)
		custHandler := http_interfaces_customer.NewCustomerHandler(custController)
		authorized.POST("/customer", custHandler.AddCustomer)
		authorized.DELETE("/customer/:id", custHandler.DeleteCustomer)
		authorized.GET("/customer/:id", custHandler.GetCustomerByID)
		authorized.GET("/customer", custHandler.GetAllCustomers)
		authorized.PUT("/customer/:id", custHandler.UpdateCustomer)
	}
}
