package repositories

import (
	"github.com/vinihss/aiqfome/internal/domain/customer"
	"github.com/vinihss/aiqfome/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(f customer.Customer) (customer.Customer, error) {
	model := models.Customer{
		Name:  f.Name,
		Email: f.Email,
	}

	if err := r.db.Create(&model).Error; err != nil {
		return customer.Customer{}, err
	}

	return customer.Customer{
		ID:    model.ID,
		Name:  model.Name,
		Email: model.Email,
	}, nil
}

func (r *CustomerRepository) FindByID(id uint) (customer.Customer, error) {
	var model models.Customer

	if err := r.db.First(&model, id).Error; err != nil {
		return customer.Customer{}, err
	}

	return customer.Customer{
		ID:    model.ID,
		Name:  model.Name,
		Email: model.Email,
	}, nil
}

func (r *CustomerRepository) Delete(id uint) error {
	var model models.Customer

	if err := r.db.First(&model, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepository) Update(new customer.Customer) (customer.Customer, error) {
	var model models.Customer

	if err := r.db.First(&model, new.ID).Error; err != nil {
		return customer.Customer{}, err
	}

	model.Email = new.Email
	model.Name = new.Name

	if err := r.db.Save(&model).Error; err != nil {
		return customer.Customer{}, err
	}

	return customer.Customer{
		ID:    model.ID,
		Name:  model.Name,
		Email: model.Email,
	}, nil
}

func (r *CustomerRepository) FindAll(offset, size int) ([]customer.Customer, error) {
	var models []models.Customer

	if err := r.db.Offset(offset).Limit(size).Find(&models).Error; err != nil {
		return nil, err
	}

	var customers []customer.Customer
	for _, model := range models {
		customers = append(customers, customer.Customer{
			ID:    model.ID,
			Name:  model.Name,
			Email: model.Email,
		})
	}

	return customers, nil
}
