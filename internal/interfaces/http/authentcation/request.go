package http_interfaces_authentication

type CreateAuthenticationRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
