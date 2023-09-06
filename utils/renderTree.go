package utils

import (
	"itsva-puestos/models"
)

func RenderTree(puestos []models.PuestoWithDetails, jefeID uint) []models.Tree {

	var tree []models.Tree
	for _, puesto := range puestos {
		if puesto.IDJefe == jefeID {
			subordinados := RenderTree(puestos, puesto.ID)
			tree = append(tree, models.Tree{
				Puesto:       puesto,
				Subordinados: subordinados,
			})
		}
	}
	return tree
}
