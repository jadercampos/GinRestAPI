package configuration

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Environment Environment
	Database    Database
}

type Environment struct {
	Name string `mapstructure:"NAME"`
}

type Database struct {
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Name     string `mapstructure:"NAME"`
	SSLMode  string `mapstructure:"SSLMODE"`
}

func LoadConfig() Config {

	// Set the file name of the configurations file
	viper.SetConfigName("configuration")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./configuration")

	// // Enable VIPER to read Environment Variables
	// viper.AutomaticEnv()

	viper.SetConfigType("yml")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("cannot read cofiguration: " + err.Error())
	}

	config := Config{}

	err = viper.Unmarshal(&config)

	if err != nil {
		log.Fatal("environment cant be loaded: ", err)
	}

	return config
}
