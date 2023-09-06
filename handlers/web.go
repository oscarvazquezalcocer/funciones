package handlers

import (
	"funciones/models"
	"funciones/services"
	"funciones/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func List(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var puestos []models.Funcion
	db.Find(&puestos)

	var nombresJefes = make(map[uint]string)
	for _, puesto := range puestos {
		if puesto.IDJefe != 0 {
			var puestoPadre models.Funcion
			db.First(&puestoPadre, puesto.IDJefe)
			nombresJefes[puesto.IDJefe] = puestoPadre.Nombre
		}
	}

	funcionesWithDetails, err := utils.GetFuncionesWithDetails(puestos)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.HTML(http.StatusOK, "list.html", gin.H{"funciones": funcionesWithDetails, "nombresJefes": nombresJefes})
}

func ShowForm(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	puestos, err := services.GetListPuestoFromAPI()

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	users, err := services.GetListUserFromAPI()

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var jefes []models.Funcion
	db.Find(&jefes)

	c.HTML(http.StatusOK, "create.html", gin.H{"jefes": jefes, "puestos": puestos, "users": users})

}

func Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var newPuesto models.Funcion
	if err := c.ShouldBind(&newPuesto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Checamos que el puesto que se reciba sea valido
	if newPuesto.IDJefe != 0 {
		var parent models.Funcion
		result := db.First(&parent, newPuesto.IDJefe)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ParentID no válido"})
			return
		}
	}

	if err := db.Create(&newPuesto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//c.JSON(http.StatusOK, newUser)
	c.Redirect(http.StatusSeeOther, "/")
}

func Show(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var puesto models.Funcion
	result := db.First(&puesto, id)
	if result.Error != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"message": "Puesto no encontrado"})
		return
	}

	puestos, err := services.GetListPuestoFromAPI()

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	users, err := services.GetListUserFromAPI()

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var jefes []models.Funcion
	db.Find(&jefes)

	funcionesWithDetails, err := utils.GetFuncionWithDetails(puesto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.HTML(http.StatusOK, "show.html", gin.H{"funcion": funcionesWithDetails, "jefes": jefes, "puestos": puestos, "users": users})
}

func Update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var puesto models.Funcion
	result := db.First(&puesto, id)
	if result.Error != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"message": "Puesto no encontrado"})
		return
	}
	if err := c.ShouldBind(&puesto); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": err.Error()})
		return
	}

	db.Save(&puesto)
	c.Redirect(http.StatusSeeOther, "/")
}

func Delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var puesto models.Funcion
	result := db.First(&puesto, id)
	if result.Error != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"message": "Puesto no encontrado"})
		return
	}
	db.Unscoped().Delete(&puesto)
	c.Redirect(http.StatusSeeOther, "/")
}

func TreeView(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var puestos []models.Funcion
	db.Find(&puestos)

	puestosWithDetails, err := utils.GetFuncionesWithDetails(puestos)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	topLevel := utils.RenderTree(puestosWithDetails, 0) // 0 representa el jefe raíz

	//c.JSON(http.StatusOK, gin.H{"tree": topLevel, "users": users})
	c.HTML(http.StatusOK, "tree.html", gin.H{"tree": topLevel})
}
