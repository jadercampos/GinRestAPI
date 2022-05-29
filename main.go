package main

import (
	"github.com/jadercampos/GinRestAPI/database"
	"github.com/jadercampos/GinRestAPI/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
