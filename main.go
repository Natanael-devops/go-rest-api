package main

import (
	"fmt"

	"github.com/Natanael-devops/go-rest-api/models"
	"github.com/Natanael-devops/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("iniciando o servidor rest com Go")
	routes.HandleRequest()
}
