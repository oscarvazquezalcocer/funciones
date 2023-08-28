package utils

import (
	"itsva-puestos/models"
)

func RenderTree(puestos []models.Puesto, users []models.UserAPI, jefeID uint) []models.Tree {

	var tree []models.Tree
	for _, puesto := range puestos {
		if puesto.IDJefe == jefeID {
			user := getUserForPuesto(users, puesto.ID)
			subordinados := RenderTree(puestos, users, puesto.ID)
			tree = append(tree, models.Tree{
				Puesto:       puesto,
				User:         user,
				Subordinados: subordinados,
			})
		}
	}
	return tree
}

func getUserForPuesto(users []models.UserAPI, puestoID uint) models.UserAPI {
	for _, user := range users {
		if user.IDPuesto == puestoID {
			return user
		}
	}
	return models.UserAPI{} // Si no se encuentra el usuario, devuelve un valor vacío
}
