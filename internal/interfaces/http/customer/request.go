package http_interfaces_customer

type CreateCustomerRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateCustomerRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
