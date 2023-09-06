package routes

import (
	"funciones/handlers"

	"github.com/gin-gonic/gin"
)

func AddWEB(r *gin.Engine) {

	r.GET("/", handlers.List)
	r.GET("/create", handlers.ShowForm)
	r.POST("/create", handlers.Create)
	r.GET("/:id", handlers.Show)
	r.POST("/:id/update", handlers.Update)
	r.POST("/:id/delete", handlers.Delete)
	r.GET("/tree", handlers.TreeView)

}
