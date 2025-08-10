package http_interfaces_favorite

import (
	"github.com/vinihss/aiqfome/internal/usecases/favorite"
)

type FavoriteController struct {
	createUC *favorite.CreateFavoriteUseCase
}

func NewFavoriteController(createUC *favorite.CreateFavoriteUseCase) *FavoriteController {
	return &FavoriteController{createUC: createUC}
}

func (ctrl *FavoriteController) CreateFavorite(req CreateFavoriteRequest) (FavoriteResponse, error) {
	fav, err := ctrl.createUC.Execute(favorite.CreateFavoriteInput{
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
	})
	if err != nil {
		return FavoriteResponse{}, err
	}

	return FavoriteResponse{
		ID:         fav.ID,
		CustomerID: fav.CustomerID,
		ProductID:  fav.ProductID,
		Product:    fav.Product,
	}, nil
}
