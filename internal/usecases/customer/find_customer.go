package customer

import (
	"github.com/vinihss/aiqfome/internal/domain/customer"
)

type FindCustomerUseCase struct {
	repo CustomerRepository
}

func NewFindCustomerUseCase(repo CustomerRepository) *FindCustomerUseCase {
	return &FindCustomerUseCase{repo: repo}
}

func (uc *FindCustomerUseCase) Execute(id uint) (customer.Customer, error) {
	entity, err := uc.repo.FindByID(id)
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func (uc *FindCustomerUseCase) ExecuteAll(page, size int) ([]customer.Customer, error) {
	offset := (page - 1) * size
	customers, err := uc.repo.FindAll(offset, size)
	if err != nil {
		return nil, err
	}
	return customers, nil
}
