package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/pranav-patidar9354/job-portal-v2/backend/pkg/services"
)

type SavedJobHandler struct{ Service *services.SavedJobService }

func (h *SavedJobHandler) Toggle(c *gin.Context) {
	jobID, err := strconv.ParseUint(c.Param("jobID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job id"})
		return
	}

	saved, err := h.Service.Toggle(uint(jobID), c.MustGet("user_id").(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"saved": saved})
}

func (h *SavedJobHandler) List(c *gin.Context) {
	saved, err := h.Service.List(c.MustGet("user_id").(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"saved_jobs": saved})
}
