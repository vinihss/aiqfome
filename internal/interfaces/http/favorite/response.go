package http_interfaces_favorite

type FavoriteResponse struct {
	ID         uint   `json:"id"`
	CustomerID uint   `json:"customer_id"`
	ProductID  uint   `json:"product_id"`
	Product    string `json:"product"`
}

func ToFavoriteResponse(id, customerID, productID uint, product string) FavoriteResponse {
	return FavoriteResponse{
		ID:         id,
		CustomerID: customerID,
		ProductID:  productID,
		Product:    product,
	}
}
