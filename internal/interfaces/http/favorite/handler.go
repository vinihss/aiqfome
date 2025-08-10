package http_interfaces_favorite

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/vinihss/aiqfome/internal/infrastructure/external_epis"
	usecase "github.com/vinihss/aiqfome/internal/usecases/favorite"
)

type FavoriteHandler struct {
	controller *FavoriteController
}

func NewFavoriteHandler(controller *FavoriteController) *FavoriteHandler {
	return &FavoriteHandler{controller: controller}
}

// Create godoc
// @Summary Adicionar produto aos favoritos do cliente
// @Tags Favorites
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param body body AddFavoriteRequest true "Produto favorito"
// @Success 201 {object} FavoriteResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customer/{id}/favorites [post]
// @Security BearerAuth
func (h *FavoriteHandler) Create(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("id"))
	if err != nil || customerID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	var req AddFavoriteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fav, err := h.controller.Add(c.Request.Context(), uint(customerID), req.ProductID)
	if err != nil {

		if err == usecase.ErrAlreadyFavorited {
			c.JSON(http.StatusConflict, gin.H{"error": "produto já favoritado para este cliente"})
			return
		}
		if err == external_epis.ErrProductNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "produto não encontrado"})
			return
		}
		msg := strings.ToLower(err.Error())
		if strings.Contains(msg, "unique") || strings.Contains(msg, "duplicate") {
			c.JSON(http.StatusConflict, gin.H{"error": "produto já favoritado para este cliente"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ToFavoriteResponse(fav.ID, fav.CustomerID, fav.ProductID, fav.Title, fav.ImageUrl, fav.Price))
}

// List godoc
// @Summary Listar favoritos do cliente
// @Tags Favorites
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {array} FavoriteResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customer/{id}/favorites [get]
// @Security BearerAuth
func (h *FavoriteHandler) List(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("id"))
	if err != nil || customerID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	items, err := h.controller.List(c.Request.Context(), uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	out := make([]FavoriteResponse, 0, len(items))
	for _, it := range items {
		out = append(out, ToFavoriteResponse(it.ID, it.CustomerID, it.ProductID, it.Title, it.ImageUrl, it.Price))
	}
	c.JSON(http.StatusOK, out)
}

// Delete godoc
// @Summary Remover produto dos favoritos do cliente
// @Tags Favorites
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param productId path int true "Product ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customer/{id}/favorites/{productId} [delete]
// @Security BearerAuth
func (h *FavoriteHandler) Delete(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("id"))
	if err != nil || customerID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	productID, err := strconv.Atoi(c.Param("productId"))
	if err != nil || productID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	if err := h.controller.Remove(c.Request.Context(), uint(customerID), uint(productID)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "favorito não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
