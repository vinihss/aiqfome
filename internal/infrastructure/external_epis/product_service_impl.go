package external_epis

import (
	"encoding/json"
	"fmt"
	"github.com/vinihss/aiqfome/internal/domain/external_services"
	"net/http"
)

type FakeStoreProductService struct {
	Url string `json:"url"`
}

func NewFakeStoreProductService() *FakeStoreProductService {
	return &FakeStoreProductService{
		Url: "https://fakestoreapi.com/products",
	}
}
func (f *FakeStoreProductService) GetProductByID(id string) (*external_services.ProductData, error) {
	url := fmt.Sprintf("%s/%s", f.Url, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		ID          int     `json:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Category    string  `json:"category"`
		Price       float32 `json:"price"`
		Image       string  `json:"image"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &external_services.ProductData{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Category:    result.Category,
		Price:       result.Price,
		Image:       result.Image,
	}, nil
}
