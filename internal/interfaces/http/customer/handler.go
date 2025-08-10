package http_interfaces_customer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type CustomerHandler struct {
	controller *CustomerController
}

func NewCustomerHandler(controller *CustomerController) *CustomerHandler {
	return &CustomerHandler{controller: controller}
}

// Create godoc
// @Summary Add customer product
// @Description Creates a favorite for a given customer and product
// @Tags Customer
// @Accept json
// @Produce json
// @Param favorite body CreateCustomerRequest true "Customer data"
// @Success 200 {object} CustomerResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customer [post]
// @Security BearerAuth
func (h *CustomerHandler) Create(c *gin.Context) {
	var req CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.controller.CreateCustomer(req)
	if err != nil {
		if isUniqueEmailErr(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary Delete customer
// @Description Deletes a customer by their ID
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customer/{id} [delete]
// @Security BearerAuth
func (h *CustomerHandler) Delete(c *gin.Context) {

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

// Update godoc
// @Summary Update customer
// @Description Updates a customer's information by their ID
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param customer body UpdateCustomerRequest true "Updated customer data"
// @Success 200 {object} CustomerResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customer/{id} [put]
// @Security BearerAuth
func (h *CustomerHandler) Update(c *gin.Context) {
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
		if isUniqueEmailErr(err) {
			c.JSON(http.StatusConflict, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// FindByID godoc
// @Summary Get customer by ID
// @Description Retrieves a customer's information by their ID
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} CustomerResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customer/{id} [get]
// @Security BearerAuth
func (h *CustomerHandler) FindByID(c *gin.Context) {
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
// @Tags Customer
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

func isUniqueEmailErr(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())

	return (strings.Contains(msg, "unique") || strings.Contains(msg, "duplicate") || strings.Contains(msg, "already exists")) &&
		strings.Contains(msg, "email")
}
