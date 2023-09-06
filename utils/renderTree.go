package utils

import "funciones/models"

func RenderTree(puestos []models.FuncionWithDetails, jefeID uint) []models.Tree {

	var tree []models.Tree
	for _, puesto := range puestos {
		if puesto.IDJefe == jefeID {
			subordinados := RenderTree(puestos, puesto.ID)
			tree = append(tree, models.Tree{
				Funcion:      puesto,
				Subordinados: subordinados,
			})
		}
	}
	return tree
}
