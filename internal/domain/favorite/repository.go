package favorite

type Repository interface {
	Create(f Favorite) (Favorite, error)
	FindByID(id string) (*Favorite, error)
}
