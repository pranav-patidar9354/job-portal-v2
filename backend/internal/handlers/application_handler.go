package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/dto"
	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/services"
)

type ApplicationHandler struct{ Service *services.ApplicationService }

func (h *ApplicationHandler) Apply(c *gin.Context) {
	jobID, err := strconv.ParseUint(c.Param("jobID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job id"})
		return
	}

	var req dto.ApplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	app, err := h.Service.Apply(uint(jobID), c.MustGet("user_id").(uint), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"application": app})
}

func (h *ApplicationHandler) Mine(c *gin.Context) {
	apps, err := h.Service.Mine(c.MustGet("user_id").(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"applications": apps})
}

func (h *ApplicationHandler) Withdraw(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid application id"})
		return
	}

	if err := h.Service.Withdraw(uint(id), c.MustGet("user_id").(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "application withdrawn"})
}

func (h *ApplicationHandler) Applicants(c *gin.Context) {
	jobID, err := strconv.ParseUint(c.Param("jobID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job id"})
		return
	}

	apps, err := h.Service.Applicants(uint(jobID), c.MustGet("user_id").(uint))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"applications": apps})
}

func (h *ApplicationHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid application id"})
		return
	}

	var req dto.UpdateApplicationStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	app, err := h.Service.UpdateStatus(uint(id), c.MustGet("user_id").(uint), req)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"application": app})
}
