package utils

import (
	"funciones/models"
	"funciones/services"
)

func GetFuncionesWithDetails(funciones []models.Funcion) ([]models.FuncionWithDetails, error) {

	var funcionWithDetails []models.FuncionWithDetails

	for _, funcion := range funciones {

		details, err := GetFuncionWithDetails(funcion)
		if err != nil {
			return nil, err
		}
		funcionWithDetails = append(funcionWithDetails, details)
	}

	return funcionWithDetails, nil
}

func GetFuncionWithDetails(funcion models.Funcion) (models.FuncionWithDetails, error) {

	puesto, err := services.GetPuestoFromIdAPI(funcion.IDPuesto)
	if err != nil {
		return models.FuncionWithDetails{}, err
	}

	user, err := services.GetUserFromIdAPI(funcion.IDUsuario)
	if err != nil {
		return models.FuncionWithDetails{}, err
	}

	funcionWithDetails := models.FuncionWithDetails{
		ID: funcion.ID,

		Nombre:        funcion.Nombre,
		IDPuesto:      funcion.IDPuesto,
		NombrePuesto:  puesto.Nombre,
		IDUsuario:     funcion.IDUsuario,
		NombreUsuario: user.Nombre,
		IDJefe:        funcion.IDJefe,
	}

	return funcionWithDetails, nil
}
