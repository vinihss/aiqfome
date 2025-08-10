package http_interfaces_customer

type CreateCustomerRequest struct {
	Name  string `json:"name"  binding:"required,min=2,max=100"`
	Email string `json:"email" binding:"required,email,max=254"`
}

type UpdateCustomerRequest struct {
	Name  string `json:"name"  binding:"omitempty,min=2,max=100"`
	Email string `json:"email" binding:"omitempty,email,max=254"`
}
