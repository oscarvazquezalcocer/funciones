package handlers

import (
	"itsva-puestos/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PuestoUserAssociation(c *gin.Context) {

	token := utils.GenerateToken()

	c.JSON(http.StatusOK, token)
}
