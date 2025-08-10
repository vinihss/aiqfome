package customer

type Repository interface {
	Create(f Customer) (Customer, error)
	FindByID(id string) (*Customer, error)
	Delete(id string) (*Customer, error)
	Update(f Customer) (*Customer, error)
}
