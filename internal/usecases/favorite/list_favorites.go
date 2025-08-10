package favorite

import (
	"context"

	domain "github.com/vinihss/aiqfome/internal/domain/favorite"
)

type ListFavoritesUseCase struct {
	repo domain.Repository
}

func NewListFavoritesUseCase(repo domain.Repository) *ListFavoritesUseCase {
	return &ListFavoritesUseCase{repo: repo}
}

func (uc *ListFavoritesUseCase) Execute(_ context.Context, customerID uint) ([]domain.Favorite, error) {
	return uc.repo.ListByCustomer(customerID)
}
