package database

import (
	"fmt"
	"log"

	"github.com/jadercampos/GinRestAPI/configuration"
	"github.com/jadercampos/GinRestAPI/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	DB, err = gorm.Open(postgres.Open(getConString()))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}

func getConString() string {
	config := configuration.LoadConfig().Database
	conString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Host, config.Username, config.Password, config.Name, config.Port, config.SSLMode)
	return conString
}
