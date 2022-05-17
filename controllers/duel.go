package controllers

import (
	"arena/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DuelController struct{}

var duel = new(models.Duel)

func (DuelController) PostDuel(c *gin.Context) {
	var gs []models.Gladiator
	if err := c.BindJSON(&gs); err != nil {
		msg := fmt.Sprint("Cannot bind request body as JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	if res, err := duel.Resolve(gs); err == nil {
		c.IndentedJSON(http.StatusOK, res)
	} else {
		msg := fmt.Sprint("Cannot resolve duel:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
}
