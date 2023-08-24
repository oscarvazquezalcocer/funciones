package utils

import (
	"itsva-puestos/models"
	"itsva-puestos/services"
)

func GetPuestosWithDetails(puestos []models.Puesto) ([]models.PuestoWithDetails, error) {

	var puestoWithDetails []models.PuestoWithDetails

	for _, puesto := range puestos {

		details, err := GetPuestoWithDetails(puesto)
		if err != nil {
			return nil, err
		}
		puestoWithDetails = append(puestoWithDetails, details)
	}

	return puestoWithDetails, nil
}

func GetPuestoWithDetails(puesto models.Puesto) (models.PuestoWithDetails, error) {

	funcion, err := services.GetFuncionFromIdAPI(puesto.IDFuncion)
	if err != nil {
		return models.PuestoWithDetails{}, err
	}

	puestoWithDetails := models.PuestoWithDetails{
		ID: puesto.ID,

		Nombre:        puesto.Nombre,
		Descripcion:   puesto.Descripcion,
		Repetible:     puesto.Repetible,
		IDJefe:        puesto.IDJefe,
		NombreFuncion: funcion.Nombre,
	}

	return puestoWithDetails, nil
}
