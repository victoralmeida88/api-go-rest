package main

import (
	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
