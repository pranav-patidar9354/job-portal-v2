package middleware

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	cfg := cors.DefaultConfig()

	origins := []string{"http://localhost:5173"}
	if envOrigins := os.Getenv("ALLOWED_ORIGINS"); envOrigins != "" {
		for _, o := range strings.Split(envOrigins, ",") {
			trimmed := strings.TrimSpace(o)
			if trimmed != "" {
				origins = append(origins, trimmed)
			}
		}
	} else if clientURL := os.Getenv("CLIENT_URL"); clientURL != "" {
		origins = append(origins, strings.TrimSpace(clientURL))
	}

	cfg.AllowOrigins = origins
	cfg.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	cfg.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	return cors.New(cfg)
}
