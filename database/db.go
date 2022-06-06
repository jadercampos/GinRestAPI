package database

import (
	"log"

	"github.com/jadercampos/GinRestAPI/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	//stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	stringDeConexao := "host=ec2-3-211-221-185.compute-1.amazonaws.com user=dxgetmqtrczbmj password=807d01583250185ecba479ccd30de6f70466c6b2033bee18d4a90565a607b8d8 dbname=da0gk0gmlam0dq port=5432 sslmode=require"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
