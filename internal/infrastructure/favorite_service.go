package infrastructure

import (
	"github.com/vinihss/aiqfome/config"
	"github.com/vinihss/aiqfome/internal/domain/favorite"
	interfaces "github.com/vinihss/aiqfome/internal/interfaces/http/favorite"
)

func CreateFavorite(dto interfaces.CreateFavoriteDTO) (favorite.Favorite, error) {
	newFavorite := favorite.Favorite{
		CustomerID: dto.UserID,
		ProductID:  dto.ProductID,
		Product:    dto.Product,
	}

	if err := config.DB.Create(&newFavorite).Error; err != nil {
		return favorite.Favorite{}, err
	}

	return newFavorite, nil
}
