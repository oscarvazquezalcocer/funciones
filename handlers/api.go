package handlers

import (
	"itsva-puestos/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListAPI(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var puestos []models.Puesto
	db.Find(&puestos)

	c.JSON(http.StatusOK, gin.H{"puestos": puestos})
}

func ShowAPI(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var puesto models.Puesto
	if err := db.First(&puesto, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Puesto no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"puesto": puesto})
}
