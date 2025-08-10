package http_interfaces_favorite

type AddFavoriteRequest struct {
	ProductID uint `json:"product_id" binding:"required,numeric,min=1"` // ProductID must be a positive integer
}
