package http_interfaces_customer

import (
	"github.com/vinihss/aiqfome/internal/usecases/customer"
)

type CustomerController struct {
	createUC *customer.CreateCustomerUseCase
	deleteUC *customer.DeleteCustomerUseCase
	findUC   *customer.FindCustomerUseCase
	updateUC *customer.UpdateCustomerUseCase
}

func NewCustomerController(
	createUC *customer.CreateCustomerUseCase,
	deleteUC *customer.DeleteCustomerUseCase,
	findUC *customer.FindCustomerUseCase,
	updateUC *customer.UpdateCustomerUseCase,
) *CustomerController {
	return &CustomerController{createUC: createUC, deleteUC: deleteUC, findUC: findUC, updateUC: updateUC}
}

func (ctrl *CustomerController) CreateCustomer(req CreateCustomerRequest) (CustomerResponse, error) {
	fav, err := ctrl.createUC.Execute(customer.CreateCustomerInput{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return CustomerResponse{}, err
	}

	return CustomerResponse{
		ID:    fav.ID,
		Name:  fav.Name,
		Email: fav.Email,
	}, nil
}

func (ctrl *CustomerController) GetCustomer(id uint) (CustomerResponse, error) {
	fav, err := ctrl.findUC.Execute(id)
	if err != nil {
		return CustomerResponse{}, err
	}

	return CustomerResponse{
		ID:    fav.ID,
		Name:  fav.Name,
		Email: fav.Email,
	}, nil
}

func (ctrl *CustomerController) DeleteCustomer(id uint) error {
	err := ctrl.deleteUC.Execute(id)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl *CustomerController) UpdateCustomer(id uint, req UpdateCustomerRequest) (CustomerResponse, error) {
	input, err := ctrl.updateUC.Execute(customer.UpdateCustomerInput{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return CustomerResponse{}, err
	}

	return CustomerResponse{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	}, nil
}

func (ctrl *CustomerController) GetAllCustomers(page, size int) ([]CustomerResponse, error) {
	customers, err := ctrl.findUC.ExecuteAll(page, size)
	if err != nil {
		return nil, err
	}

	var responses []CustomerResponse
	for _, c := range customers {
		responses = append(responses, CustomerResponse{
			ID:    c.ID,
			Name:  c.Name,
			Email: c.Email,
		})
	}

	return responses, nil
}
