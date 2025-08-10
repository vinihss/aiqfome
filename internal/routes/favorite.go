package routes

import (
	"github.com/gin-gonic/gin"

	httpfav "github.com/vinihss/aiqfome/internal/interfaces/http/favorite"
)

func RegisterFavoriteRoutes(r *gin.Engine, handler *httpfav.FavoriteHandler) {

	r.POST("/customer/:id/favorites", handler.Create)
	r.GET("/customer/:id/favorites", handler.List)
	r.DELETE("/customer/:id/favorites/:productId", handler.Delete)

}
