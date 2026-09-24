package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/pranav-patidar9354/job-portal-v2/backend/pkg/dto"
	"github.com/pranav-patidar9354/job-portal-v2/backend/pkg/services"
)

type JobHandler struct{ Service *services.JobService }

func (h *JobHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "9"))
	if page < 1 { page = 1 }
	if limit < 1 { limit = 9 }
	if limit > 50 { limit = 50 }

	minSalary, _ := strconv.ParseFloat(c.DefaultQuery("min_salary", "0"), 64)
	maxSalary, _ := strconv.ParseFloat(c.DefaultQuery("max_salary", "0"), 64)

	q := dto.JobQuery{
		Search: c.Query("search"),
		Location: c.Query("location"),
		Company: c.Query("company"),
		JobType: c.Query("job_type"),
		MinSalary: minSalary,
		MaxSalary: maxSalary,
		Sort: c.DefaultQuery("sort", "newest"),
		Page: page,
		Limit: limit,
	}

	jobs, total, err := h.Service.List(q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))
	c.JSON(http.StatusOK, gin.H{
		"jobs": jobs,
		"pagination": gin.H{
			"page": page, "limit": limit,
			"total": total, "total_pages": totalPages,
		},
	})
}

func (h *JobHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job id"})
		return
	}

	job, err := h.Service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "job not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"job": job})
}

func (h *JobHandler) Create(c *gin.Context) {
	var req dto.CreateJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job, err := h.Service.Create(req, c.MustGet("user_id").(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"job": job})
}

func (h *JobHandler) OwnJobs(c *gin.Context) {
	jobs, err := h.Service.OwnJobs(c.MustGet("user_id").(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"jobs": jobs})
}

func (h *JobHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job id"})
		return
	}

	var req dto.UpdateJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job, err := h.Service.Update(uint(id), c.MustGet("user_id").(uint), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"job": job})
}

func (h *JobHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job id"})
		return
	}

	if err := h.Service.Delete(uint(id), c.MustGet("user_id").(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "job deleted"})
}
