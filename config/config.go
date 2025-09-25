package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("⚠️ Aucune config trouvée, valeurs par défaut utilisées")
	} else {
		fmt.Println("✅ Config chargée :", viper.AllSettings())
	}
}
