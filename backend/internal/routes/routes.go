package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pranav-patidar9354/job-portal-v2/backend/config"
	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/handlers"
	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/middleware"
	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/repositories"
	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/services"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())

	userRepo := &repositories.UserRepository{DB: config.DB}
	jobRepo := &repositories.JobRepository{DB: config.DB}
	appRepo := &repositories.ApplicationRepository{DB: config.DB}
	savedRepo := &repositories.SavedJobRepository{DB: config.DB}

	authService := &services.AuthService{Users: userRepo}
	jobService := &services.JobService{Jobs: jobRepo}
	appService := &services.ApplicationService{Apps: appRepo, Jobs: jobRepo}
	savedService := &services.SavedJobService{Saved: savedRepo, Jobs: jobRepo}

	authHandler := &handlers.AuthHandler{Service: authService}
	jobHandler := &handlers.JobHandler{Service: jobService}
	appHandler := &handlers.ApplicationHandler{Service: appService}
	savedHandler := &handlers.SavedJobHandler{Service: savedService}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Job Portal V2 API is running"})
	})

	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
	auth.GET("/me", middleware.AuthMiddleware(), authHandler.Me)

	api.GET("/jobs", jobHandler.List)
	api.GET("/jobs/:id", jobHandler.Get)

	recruiter := api.Group("/recruiter")
	recruiter.Use(middleware.AuthMiddleware(), middleware.RequireRole("recruiter"))
	recruiter.POST("/jobs", jobHandler.Create)
	recruiter.GET("/jobs", jobHandler.OwnJobs)
	recruiter.PUT("/jobs/:id", jobHandler.Update)
	recruiter.DELETE("/jobs/:id", jobHandler.Delete)
	recruiter.GET("/jobs/:jobID/applicants", appHandler.Applicants)
	recruiter.PATCH("/applications/:id/status", appHandler.UpdateStatus)

	candidate := api.Group("/candidate")
	candidate.Use(middleware.AuthMiddleware(), middleware.RequireRole("candidate"))
	candidate.POST("/jobs/:jobID/apply", appHandler.Apply)
	candidate.GET("/applications", appHandler.Mine)
	candidate.DELETE("/applications/:id", appHandler.Withdraw)
	candidate.POST("/jobs/:jobID/save", savedHandler.Toggle)
	candidate.GET("/saved-jobs", savedHandler.List)

	return r
}
