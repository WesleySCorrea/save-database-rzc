package main

import (
	"saveDatabase/project/configs/dataBase"
	"saveDatabase/project/configs/server"
)

func main() {
	// Inicie Conecção com Banco de dados
	db := dataBase.Conn()

	// Feche Conecção com Banco de dados
	dataBase.CloseConnection(db)

	// Inicie o servidor
	server.StartServer()
}
