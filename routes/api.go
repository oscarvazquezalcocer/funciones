package routes

import (
	"funciones/handlers"

	"github.com/gin-gonic/gin"
)

func AddAPI(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.GET("/", handlers.ListAPI)
		api.GET("/:id", handlers.ShowAPI)
		api.GET("/tree", handlers.TreeAPI)

	}

}
