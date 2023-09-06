package handlers

import (
	"funciones/models"
	"funciones/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListAPI(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var puestos []models.Funcion
	db.Find(&puestos)

	c.JSON(http.StatusOK, gin.H{"puestos": puestos})
}

func ShowAPI(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var puesto models.Funcion
	if err := db.First(&puesto, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Puesto no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"puesto": puesto})
}

func TreeAPI(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var puestos []models.Funcion
	db.Find(&puestos)

	puestosWithDetails, err := utils.GetFuncionesWithDetails(puestos)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	topLevel := utils.RenderTree(puestosWithDetails, 0) // 0 representa el jefe raíz

	c.JSON(http.StatusOK, gin.H{"tree": topLevel})
}
