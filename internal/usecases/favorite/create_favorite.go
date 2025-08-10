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
		Product:    "Nome do Produto", // poderia vir de outro serviço/API
	}
	return uc.repo.Create(fav)
}
