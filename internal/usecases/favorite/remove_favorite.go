package favorite

import (
	"context"

	domain "github.com/vinihss/aiqfome/internal/domain/favorite"
)

type RemoveFavoriteUseCase struct {
	repo domain.Repository
}

func NewRemoveFavoriteUseCase(repo domain.Repository) *RemoveFavoriteUseCase {
	return &RemoveFavoriteUseCase{repo: repo}
}

func (uc *RemoveFavoriteUseCase) Execute(_ context.Context, customerID uint, productID uint) error {
	return uc.repo.Delete(customerID, productID)
}
