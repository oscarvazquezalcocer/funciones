package utils

import (
	"itsva-puestos/models"
)

func RenderTree(puestos []models.Puesto, jefeID uint) []models.PuestoTree {
	var tree []models.PuestoTree
	for _, puesto := range puestos {
		if puesto.IDJefe == jefeID {
			subordinados := RenderTree(puestos, puesto.ID)
			tree = append(tree, models.PuestoTree{
				Puesto:       puesto,
				Subordinados: subordinados,
			})
		}
	}
	return tree
}
