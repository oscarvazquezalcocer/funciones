package services

import (
	"encoding/json"
	"funciones/models"

	"net/http"
	"strconv"

	"github.com/spf13/viper"
)

func GetListUserFromAPI() ([]models.User, error) {

	apiUrl := viper.GetString("api-users.url")
	apiPort := viper.GetString("api-users.port")
	apiEndpoint := viper.GetString("api-users.endpoint")

	resp, err := http.Get(apiUrl + apiPort + apiEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var responseData struct {
		Users []models.User `json:"users"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, err
	}
	return responseData.Users, nil
}

func GetUserFromIdAPI(id_user uint) (models.User, error) {

	apiUrl := viper.GetString("api-users.url")
	apiPort := viper.GetString("api-users.port")
	apiEndpoint := viper.GetString("api-users.endpoint")

	id_user_string := strconv.Itoa(int(id_user))

	resp, err := http.Get(apiUrl + apiPort + apiEndpoint + id_user_string)
	if err != nil {
		return models.User{}, err
	}
	defer resp.Body.Close()

	var responseData struct {
		User models.User `json:"user"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return models.User{}, err
	}

	return responseData.User, nil
}
