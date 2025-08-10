package favorite

import (
	"github.com/vinihss/aiqfome/internal/domain/favorite"
)

type FavoriteRepository interface {
	Create(fav favorite.Favorite) (favorite.Favorite, error)
}

type CreateFavoriteInput struct {
	CustomerID uint
	ProductID  uint
	Title      string
	ImageUrl   string
	Price      float32
}

type CreateFavoriteUseCase struct {
	repo FavoriteRepository
}

func NewCreateFavoriteUseCase(repo FavoriteRepository) *CreateFavoriteUseCase {
	return &CreateFavoriteUseCase{repo: repo}
}

func (uc *CreateFavoriteUseCase) Execute(input CreateFavoriteInput) (favorite.Favorite, error) {
	fav := favorite.Favorite{
		CustomerID: input.CustomerID,
		ProductID:  input.ProductID,
		Title:      input.Title,
		ImageUrl:   input.ImageUrl,
		Price:      input.Price,
	}
	return uc.repo.Create(fav)
}
