package http_interfaces_favorite

type CreateFavoriteDTO struct {
	UserID    string `json:"user_id" binding:"required"`
	ProductID string `json:"product_id" binding:"required"`
	Product   string `json:"product" binding:"required"`
}
