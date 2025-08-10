package routes

import (
	"github.com/gin-gonic/gin"

	httpfav "github.com/vinihss/aiqfome/internal/interfaces/http/customer"
)

func RegisterCustomerRoutes(r *gin.Engine, handler *httpfav.CustomerHandler) {
	customerGroup := r.Group("/customer")
	{
		customerGroup.POST("/", handler.Create)
		customerGroup.GET("/:id", handler.FindByID)
		customerGroup.GET("/", handler.GetAllCustomers)
		customerGroup.PUT("/:id", handler.Update)
		customerGroup.DELETE("/:id", handler.Delete)
	}
}
