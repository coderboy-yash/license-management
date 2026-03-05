package handler

import (
	"net/http"

	"license-service/models"
	"license-service/service"

	"github.com/gin-gonic/gin"
)

type LicenseHandler struct {
	service *service.LicenseService
}

func NewLicenseHandler(s *service.LicenseService) *LicenseHandler {
	return &LicenseHandler{service: s}
}

func (h *LicenseHandler) CreateLicense(c *gin.Context) {

	var req models.CreateLicenseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	licenseID, err := h.service.CreateLicense(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"license_id": licenseID,
	})
}

func (h *LicenseHandler) GetLicenseFeatures(c *gin.Context) {

	licenseID := c.Param("license_id")

	features, err := h.service.GetLicenseFeatures(c, licenseID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"features": features,
	})
}
