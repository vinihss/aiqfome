package http_interfaces_favorite

type FavoriteResponse struct {
	ID         uint    `json:"id"`
	CustomerID uint    `json:"customer_id"`
	ProductID  uint    `json:"product_id"`
	Product    string  `json:"product"`
	Title      string  `json:"title"`
	ImageUrl   string  `json:"image_url"`
	Price      float32 `json:"price"`
}

func ToFavoriteResponse(id, customerID, productID uint, title string, image string, price float32) FavoriteResponse {
	return FavoriteResponse{
		ID:         id,
		CustomerID: customerID,
		ProductID:  productID,
		Title:      title,
		ImageUrl:   image,
		Price:      price,
	}
}
