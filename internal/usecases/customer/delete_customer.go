package customer

type DeleteCustomerUseCase struct {
	repo CustomerRepository
}

func NewDeleteCustomerUseCase(repo CustomerRepository) *DeleteCustomerUseCase {
	return &DeleteCustomerUseCase{repo: repo}
}

func (uc *DeleteCustomerUseCase) Execute(id uint) error {
	entity, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}
	return uc.repo.Delete(entity.ID)
}
