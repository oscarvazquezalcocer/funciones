package services

import (
	"encoding/json"
	"funciones/models"

	"net/http"
	"strconv"

	"github.com/spf13/viper"
)

func GetListPuestoFromAPI() ([]models.PuestoAPI, error) {

	apiUrl := viper.GetString("api-puestos.url")
	apiPort := viper.GetString("api-puestos.port")
	apiEndpoint := viper.GetString("api-puestos.endpoint")

	resp, err := http.Get(apiUrl + apiPort + apiEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var responseData struct {
		Puestos []models.PuestoAPI `json:"puestos"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, err
	}
	return responseData.Puestos, nil
}

func GetPuestoFromIdAPI(id_puesto uint) (models.PuestoAPI, error) {

	apiUrl := viper.GetString("api-puestos.url")
	apiPort := viper.GetString("api-puestos.port")
	apiEndpoint := viper.GetString("api-puestos.endpoint")

	id_puesto_string := strconv.Itoa(int(id_puesto))

	resp, err := http.Get(apiUrl + apiPort + apiEndpoint + id_puesto_string)
	if err != nil {
		return models.PuestoAPI{}, err
	}
	defer resp.Body.Close()

	var responseData struct {
		Puesto models.PuestoAPI `json:"puesto"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return models.PuestoAPI{}, err
	}

	return responseData.Puesto, nil
}
