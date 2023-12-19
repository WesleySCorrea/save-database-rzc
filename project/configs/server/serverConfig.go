package server

import (
	"log"
	"os"
	"saveDatabase/project/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StartServer() {
	// Carregue as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
	port := os.Getenv("API_PORT")

	// Configuração do servidor
	r := gin.Default() // Crie uma instância do Gin

	// Defina suas rotas aqui
	routes.SetupRoutes(r)

	err = r.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
