package main

import (
	"APIgin/database"
	"APIgin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
