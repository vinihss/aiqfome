package http_interfaces_favorite

import (
	"context"

	domain "github.com/vinihss/aiqfome/internal/domain/favorite"
	usecase "github.com/vinihss/aiqfome/internal/usecases/favorite"
)

type FavoriteController struct {
	addUC    *usecase.AddFavoriteUseCase
	listUC   *usecase.ListFavoritesUseCase
	removeUC *usecase.RemoveFavoriteUseCase
}

func NewFavoriteController(
	add *usecase.AddFavoriteUseCase,
	list *usecase.ListFavoritesUseCase,
	remove *usecase.RemoveFavoriteUseCase) *FavoriteController {
	return &FavoriteController{addUC: add, listUC: list, removeUC: remove}
}

func (c *FavoriteController) Add(ctx context.Context, customerID, productID uint) (domain.Favorite, error) {
	return c.addUC.Execute(ctx, customerID, productID)
}

func (c *FavoriteController) List(ctx context.Context, customerID uint) ([]domain.Favorite, error) {
	return c.listUC.Execute(ctx, customerID)
}

func (c *FavoriteController) Remove(ctx context.Context, customerID, productID uint) error {
	return c.removeUC.Execute(ctx, customerID, productID)
}
