package handlers

import (
	"itsva-puestos/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListAPIPuestos(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var puestos []models.Puesto
	db.Find(&puestos)

	c.JSON(http.StatusOK, gin.H{"puestos": puestos})
}
