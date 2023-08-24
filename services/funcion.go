package services

import (
	"encoding/json"
	"itsva-puestos/models"
	"net/http"
	"strconv"

	"github.com/spf13/viper"
)

func GetListFuncionFromAPI() ([]models.FuncionAPI, error) {

	apiUrl := viper.GetString("api-funciones.url")
	apiPort := viper.GetString("api-funciones.port")
	apiEndpoint := viper.GetString("api-funciones.endpoint")

	resp, err := http.Get(apiUrl + apiPort + apiEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var responseData struct {
		Funciones []models.FuncionAPI `json:"funciones"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, err
	}
	return responseData.Funciones, nil
}

func GetFuncionFromIdAPI(id_funcion uint) (models.FuncionAPI, error) {

	apiUrl := viper.GetString("api-funciones.url")
	apiPort := viper.GetString("api-funciones.port")
	apiEndpoint := viper.GetString("api-funciones.endpoint")

	id_funcion_string := strconv.Itoa(int(id_funcion))

	resp, err := http.Get(apiUrl + apiPort + apiEndpoint + id_funcion_string)
	if err != nil {
		return models.FuncionAPI{}, err
	}
	defer resp.Body.Close()

	var responseData struct {
		Funcion models.FuncionAPI `json:"funcion"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return models.FuncionAPI{}, err
	}

	return responseData.Funcion, nil
}
