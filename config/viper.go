package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func ConfigureViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error al leer el archivo de configuración:", err)
		return
	}
}
