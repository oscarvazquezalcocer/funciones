package main

import (
	"itsva-puestos/models"
	"itsva-puestos/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := gorm.Open(sqlite.Open(viper.GetString("db")), &gorm.Config{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en la conexion a la base de datos"})
			c.Abort()
			return
		}
		c.Set("db", db)
		c.Next()
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}
}

func main() {
	db, err := gorm.Open(sqlite.Open(viper.GetString("db")), &gorm.Config{})
	if err != nil {
		panic("Error conexión a la base de datos")
	}

	// Realizar las migraciones antes de iniciar el servidor
	db.AutoMigrate(&models.Puesto{})

	//Agregamos al Director antes que nada
	initSquemaPuestos(db)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Use(DatabaseMiddleware())
	routes.AddWEB(r)
	routes.AddAPI(r)

	r.Run(viper.GetString("PORT"))
}

func initSquemaPuestos(db *gorm.DB) {

	// Crear el primer puesto raíz
	rootPuesto := models.Puesto{
		Nombre:      "Director",
		Descripcion: "Encargado de Adminisitrar el Instituto Tecnologico Nacional Campus Valladolid",
		IDJefe:      0,
	}

	db.Create(&rootPuesto)

}
