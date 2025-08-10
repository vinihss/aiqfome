package customer

import (
	"github.com/vinihss/aiqfome/internal/domain/customer"
)

type CustomerRepository interface {
	Create(entity customer.Customer) (customer.Customer, error)
	Delete(id uint) error
	FindByID(id uint) (customer.Customer, error)
	Update(entity customer.Customer) (customer.Customer, error)
	FindAll(int, size int) ([]customer.Customer, error)
}

type CreateCustomerInput struct {
	Name  string
	Email string
}

type CreateCustomerUseCase struct {
	repo CustomerRepository
}

func NewCreateCustomerUseCase(repo CustomerRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{repo: repo}
}

func (uc *CreateCustomerUseCase) Execute(input CreateCustomerInput) (customer.Customer, error) {
	fav := customer.Customer{
		Name:  input.Name,
		Email: input.Email,
	}
	return uc.repo.Create(fav)
}
