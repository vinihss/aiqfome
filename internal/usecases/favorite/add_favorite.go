package favorite

import (
	"context"
	"errors"

	domain "github.com/vinihss/aiqfome/internal/domain/favorite"
	"github.com/vinihss/aiqfome/internal/infrastructure/external_epis"
)

var (
	ErrAlreadyFavorited = errors.New("produto já está nos favoritos do cliente")
)

type AddFavoriteUseCase struct {
	repo   domain.Repository
	client external_epis.ProductClient
}

func NewAddFavoriteUseCase(repo domain.Repository, client external_epis.ProductClient) *AddFavoriteUseCase {
	return &AddFavoriteUseCase{repo: repo, client: client}
}

func (uc *AddFavoriteUseCase) Execute(ctx context.Context, customerID uint, productID uint) (domain.Favorite, error) {

	product, err := uc.client.GetProduct(ctx, productID)
	if err != nil {
		return domain.Favorite{}, err
	}

	exists, err := uc.repo.Exists(customerID, productID)
	if err != nil {
		return domain.Favorite{}, err
	}
	if exists {
		return domain.Favorite{}, ErrAlreadyFavorited
	}

	fav := domain.Favorite{
		CustomerID: customerID,
		ProductID:  productID,
		Title:      product.Title,
		ImageUrl:   product.Image,
		Price:      product.Price,
	}

	return uc.repo.Create(fav)
}
