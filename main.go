package main

import (
	"fmt"

	"github.com/Natanael-devops/go-rest-api/database"
	"github.com/Natanael-devops/go-rest-api/models"
	"github.com/Natanael-devops/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancodeDados()
	fmt.Println("iniciando o servidor rest com Go")
	routes.HandleRequest()
}
