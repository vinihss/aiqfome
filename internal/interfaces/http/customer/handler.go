package http_interfaces_customer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	controller *CustomerController
}

func NewCustomerHandler(controller *CustomerController) *CustomerHandler {
	return &CustomerHandler{controller: controller}
}

// AddCustomer godoc
// @Summary Add customer product
// @Description Creates a favorite for a given customer and product
// @Tags customer
// @Accept json
// @Produce json
// @Param favorite body CreateCustomerRequest true "Customer data"
// @Success 200 {object} CustomerResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customer [post]
// @Security BearerAuth
func (h *CustomerHandler) AddCustomer(c *gin.Context) {
	var req CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.controller.CreateCustomer(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteCustomer godoc
// @Summary Delete customer
// @Description Deletes a customer by their ID
// @Tags customer
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customer/{id} [delete]
// @Security BearerAuth
func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = h.controller.DeleteCustomer(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}

// UpdateCustomer godoc
// @Summary Update customer
// @Description Updates a customer's information by their ID
// @Tags customer
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param customer body UpdateCustomerRequest true "Updated customer data"
// @Success 200 {object} CustomerResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customer/{id} [put]
// @Security BearerAuth
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req UpdateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.controller.UpdateCustomer(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetCustomerByID godoc
// @Summary Get customer by ID
// @Description Retrieves a customer's information by their ID
// @Tags customer
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} CustomerResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customer/{id} [get]
// @Security BearerAuth
func (h *CustomerHandler) GetCustomerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	res, err := h.controller.GetCustomer(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllCustomers godoc
// @Summary Get all customers
// @Description Retrieves a paginated list of customers
// @Tags customer
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param size query int false "Page size"
// @Success 200 {array} CustomerResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customers [get]
// @Security BearerAuth
func (h *CustomerHandler) GetAllCustomers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	res, err := h.controller.GetAllCustomers(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
