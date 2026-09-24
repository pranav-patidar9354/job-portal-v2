package handler

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/pranav-patidar9354/job-portal-v2/backend/config"
	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/routes"
)

var (
	app  *gin.Engine
	once sync.Once
)

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(func() {
		gin.SetMode(gin.ReleaseMode)
		if err := config.ConnectDatabase(); err != nil {
			log.Printf("Database connection failed: %v", err)
		}
		app = routes.SetupRouter()
	})

	if app != nil {
		app.ServeHTTP(w, r)
	} else {
		http.Error(w, "Service Unavailable: Database connection failed", http.StatusServiceUnavailable)
	}
}
