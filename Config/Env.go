package Config

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
}

func loadEnv() *Env {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() // Automatically load from environment variables if set
	err := viper.ReadInConfig()
	if err != nil {
		return nil
	} // Read .env file

	// Create an Env instance and map Viper configurations to it
	var config Env
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &config
}
