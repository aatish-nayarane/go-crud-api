package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"vm-crud/internal/model"
	"vm-crud/internal/service"
)

type VMHandler struct {
	service *service.VMService
}

func NewVMHandler(service *service.VMService) *VMHandler {
	return &VMHandler{
		service: service,
	}
}

func (h *VMHandler) CreateVM(c *gin.Context) {
	var req model.CreateVMRequest
	// Then Gin converts JSON into our request model:
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("Error binding JSON:", err)
		return
	}
	vm, err := h.service.CreateVM(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("whie creatimng vm handler :", err)
		return
	}

	c.JSON(http.StatusOK, vm)
}

func (h *VMHandler) ListVMs(c *gin.Context) {
	vms, err := h.service.ListVMs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("Error listing VMs in handler:", err)
		return
	}

	c.JSON(http.StatusOK, vms)
}

func (h *VMHandler) GetVMByID(c *gin.Context) {
	idParam := c.Param("id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		fmt.Println("Error parsing ID parameter:", err)
		return
	}

	vm, err := h.service.GetVMByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("Error getting VM by ID in handler:", err)
		return
	}

	c.JSON(http.StatusOK, vm)
}

func (h *VMHandler) UpdateVM(c *gin.Context) {
	idParam := c.Param("id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		fmt.Println("Error parsing ID parameter:", err)
		return
	}

	var req model.UpdateVMRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("Error binding JSON:", err)
		return
	}
	
	vm, err := h.service.UpdateVM(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("Error updating VM in handler:", err)
		return
	}

	c.JSON(http.StatusOK, vm)
}

func (h *VMHandler) DeleteVM(c *gin.Context) {
	idParam := c.Param("id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		fmt.Println("Error parsing ID parameter:", err)
		return
	}

	err = h.service.DeleteVM(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("Error deleting VM in handler:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "VM deleted successfully"})
}