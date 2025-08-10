package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vinihss/aiqfome/internal/infrastructure/database/repositories"
	"github.com/vinihss/aiqfome/internal/infrastructure/external_epis"
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
	db, _ := gorm.Open(postgres.Open("host=postgres user=postgres password=postgres dbname=favorites port=5432 sslmode=disable"), &gorm.Config{})

	authorized := router.Group("/")
	authorized.Use(middlewares.JWTAuth())
	{
		productClient := external_epis.NewFakeStoreClient()
		rdb := redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		})
		favRepo := repositories.NewFavoriteRepository(db, rdb)
		createFavoriteUC := favoriteuse.NewAddFavoriteUseCase(favRepo, productClient)
		listFavoriteUC := favoriteuse.NewListFavoritesUseCase(favRepo)
		removeFavoriteUC := favoriteuse.NewRemoveFavoriteUseCase(favRepo)
		favController := http_interfaces_favorite.NewFavoriteController(createFavoriteUC, listFavoriteUC, removeFavoriteUC)
		favHandler := http_interfaces_favorite.NewFavoriteHandler(favController)
		RegisterFavoriteRoutes(router, favHandler)

		custRepo := repositories.NewCustomerRepository(db)
		createCustomerUC := customeruse.NewCreateCustomerUseCase(custRepo)
		deleteCustomerUC := customeruse.NewDeleteCustomerUseCase(custRepo)
		findCustomerUC := customeruse.NewFindCustomerUseCase(custRepo)
		updateCustomerUC := customeruse.NewUpdateCustomerUseCase(custRepo)

		custController := http_interfaces_customer.NewCustomerController(createCustomerUC, deleteCustomerUC, findCustomerUC, updateCustomerUC)
		custHandler := http_interfaces_customer.NewCustomerHandler(custController)
		RegisterCustomerRoutes(router, custHandler)

	}

}
