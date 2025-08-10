package http_interfaces_favorite

type CreateFavoriteRequest struct {
	CustomerID string `json:"customer_id" binding:"required"`
	ProductID  string `json:"product_id" binding:"required"`
}
