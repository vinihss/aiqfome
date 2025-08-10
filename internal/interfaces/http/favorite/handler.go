package http_interfaces_favorite

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	controller *FavoriteController
}

func NewFavoriteHandler(controller *FavoriteController) *FavoriteHandler {
	return &FavoriteHandler{controller: controller}
}

// AddFavorite godoc
// @Summary Add favorite product
// @Description Creates a favorite for a given customer and product
// @Tags favorites
// @Accept json
// @Produce json
// @Param favorite body CreateFavoriteRequest true "Favorite data"
// @Success 200 {object} FavoriteResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /favorites [post]
// @Security BearerAuth
func (h *FavoriteHandler) AddFavorite(c *gin.Context) {
	var req CreateFavoriteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.controller.CreateFavorite(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
