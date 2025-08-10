package favorite

type Repository interface {
	Create(f Favorite) (Favorite, error)
	Exists(customerID uint, productID uint) (bool, error)
	ListByCustomer(customerID uint) ([]Favorite, error)
	Delete(customerID uint, productID uint) error
}
