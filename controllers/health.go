package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
