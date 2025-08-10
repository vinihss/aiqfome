package external_services

type ProductData struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	Image       string  `json:"image"`
}

type ProductService interface {
	GetProductByID(id string) (*ProductData, error)
	ListProducts() ([]ProductData, error)
}
