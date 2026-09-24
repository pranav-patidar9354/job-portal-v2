package middleware

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	cfg := cors.DefaultConfig()

	cfg.AllowOriginFunc = func(origin string) bool {
		// Allow localhost origins for local development
		if strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "http://127.0.0.1:") {
			return true
		}
		// Allow all Vercel deployments (*.vercel.app)
		if strings.HasSuffix(origin, ".vercel.app") {
			return true
		}
		// Allow custom specified domains
		if custom := os.Getenv("CLIENT_URL"); custom != "" && strings.EqualFold(origin, custom) {
			return true
		}
		if envOrigins := os.Getenv("ALLOWED_ORIGINS"); envOrigins != "" {
			for _, o := range strings.Split(envOrigins, ",") {
				if strings.EqualFold(origin, strings.TrimSpace(o)) {
					return true
				}
			}
		}
		return false
	}

	cfg.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	cfg.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "Access-Control-Request-Private-Network"}
	cfg.AllowCredentials = true

	corsMiddleware := cors.New(cfg)

	return func(c *gin.Context) {
		if c.Request.Header.Get("Access-Control-Request-Private-Network") == "true" {
			c.Header("Access-Control-Allow-Private-Network", "true")
		}
		corsMiddleware(c)
	}
}

