package server

import (
	"os"
	"strings"

	"github.com/colosseum-project/app-arena/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	if proxies := os.Getenv("GIN_PROXIES"); proxies != "" {
		p := strings.Split(proxies, ",")
		router.SetTrustedProxies(p)
	}

	health := new(controllers.HealthController)
	duel := new(controllers.DuelController)

	router.GET("/health", health.Status)
	router.POST("/duel", duel.PostDuel)

	return router
}
