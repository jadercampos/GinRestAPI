package main

import (
	"github.com/jadercampos/GinRestAPI/database"
	"github.com/jadercampos/GinRestAPI/models"
	"github.com/jadercampos/GinRestAPI/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Pedro", CPF: "0000000000", RG: "99999999"},
		{Nome: "João", CPF: "1111111111", RG: "88888888"},
	}
	routes.HandleRequests()
}
