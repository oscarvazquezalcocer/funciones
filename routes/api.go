package routes

import (
	"itsva-puestos/handlers"

	"github.com/gin-gonic/gin"
)

func AddAPI(r *gin.Engine) {

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/", handlers.PuestoUserAssociation)

	}

}
