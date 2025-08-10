package http_interfaces_favorite

type FavoriteResponse struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	ProductID  string `json:"product_id"`
	Product    string `json:"product"`
}
