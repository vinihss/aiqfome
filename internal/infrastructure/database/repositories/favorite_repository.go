package repositories

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"time"

	domain "github.com/vinihss/aiqfome/internal/domain/favorite"
	"github.com/vinihss/aiqfome/internal/infrastructure/database/models"
	"gorm.io/gorm"
)

type FavoriteRepository struct {
	db    *gorm.DB
	cache *redis.Client
}

const (
	// Cache TTL para lista de favoritos por cliente
	favoritesListTTL = 15 * time.Minute
	// Chave de cache para lista de favoritos
	favoritesListKey = "favorites:customer:%d"
)

func NewFavoriteRepository(db *gorm.DB, cache *redis.Client) *FavoriteRepository {
	return &FavoriteRepository{
		db:    db,
		cache: cache,
	}
}

func (r *FavoriteRepository) Create(f domain.Favorite) (domain.Favorite, error) {
	model := models.Favorite{
		CustomerID: f.CustomerID,
		ProductID:  f.ProductID,
		Title:      f.Title,
		ImageUrl:   f.ImageUrl,
		Price:      f.Price,
	}

	if err := r.db.Create(&model).Error; err != nil {
		return domain.Favorite{}, err
	}

	return domain.Favorite{
		ID:         model.ID,
		CustomerID: model.CustomerID,
		ProductID:  model.ProductID,
		Title:      model.Title,
		ImageUrl:   model.ImageUrl,
		Price:      model.Price,
	}, nil
}

func (r *FavoriteRepository) Exists(customerID uint, productID uint) (bool, error) {
	var count int64
	if err := r.db.Model(&models.Favorite{}).
		Where("customer_id = ? AND product_id = ?", customerID, productID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *FavoriteRepository) ListByCustomer(customerID uint) ([]domain.Favorite, error) {
	var rows []models.Favorite
	if err := r.db.Where("customer_id = ?", customerID).Order("id DESC").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]domain.Favorite, 0, len(rows))
	for _, m := range rows {
		out = append(out, domain.Favorite{
			ID:         m.ID,
			CustomerID: m.CustomerID,
			ProductID:  m.ProductID,
			Title:      m.Title,
			ImageUrl:   m.ImageUrl,
			Price:      m.Price,
		})
	}
	return out, nil
}

func (r *FavoriteRepository) Delete(customerID uint, productID uint) error {
	res := r.db.Where("customer_id = ? AND product_id = ?", customerID, productID).Delete(&models.Favorite{})
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Helper para traduzir erros de unicidade, caso queira tratar em camadas superiores
func isUniqueConstraintErr(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, gorm.ErrDuplicatedKey)
}
