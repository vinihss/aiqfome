package repositories

import (
	"github.com/vinihss/aiqfome/internal/domain/favorite"
	"github.com/vinihss/aiqfome/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type FavoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) *FavoriteRepository {
	return &FavoriteRepository{db: db}
}

func (r *FavoriteRepository) Create(f favorite.Favorite) (favorite.Favorite, error) {
	model := models.Favorite{
		ID:         f.ID,
		CustomerID: f.CustomerID,
		ProductID:  f.ProductID,
		Product:    f.Product,
	}

	if err := r.db.Create(&model).Error; err != nil {
		return favorite.Favorite{}, err
	}

	// converte de volta para entidade de domínio
	return favorite.Favorite{
		ID:         model.ID,
		CustomerID: model.CustomerID,
		ProductID:  model.ProductID,
		Product:    model.Product,
	}, nil
}
