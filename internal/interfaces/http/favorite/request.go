package http_interfaces_favorite

type AddFavoriteRequest struct {
	ProductID uint `json:"product_id" binding:"required,gt=0"`
}
