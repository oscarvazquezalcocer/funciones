package routes

import (
	"itsva-puestos/handlers"

	"github.com/gin-gonic/gin"
)

func AddWEB(r *gin.Engine) {

	r.GET("/", handlers.ListPuestos)
	r.GET("/create", handlers.ShowFormPuesto)
	r.POST("/create", handlers.CreatePuesto)
	r.GET("/:id", handlers.ViewPuesto)
	r.POST("/:id/update", handlers.UpdatePuesto)
	r.POST("/:id/delete", handlers.DeletePuesto)
	r.GET("/tree", handlers.TreeView)

}
