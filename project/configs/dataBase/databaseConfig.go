package dataBase

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dBConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func GetDBConfig() (*dBConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Erro ao carregar o arquivo .env: %v", err)
	}

	config := &dBConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}

	if err != nil {
		log.Fatal("Erro ao obter a configuração do banco de dados: ", err)
	}
	return config, nil
}

func CreateConnection(config dBConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao criar a conexão com o banco de dados: ", err)
	}

	fmt.Println("Conexão estabelecida com sucesso!")
	return db, nil
}

func Conn() *gorm.DB {
	dbConfig, err := GetDBConfig()
	if err != nil {
		fmt.Println("Problemas ao carregar configuração do Banco de Dados!")
		return nil
	}

	db, err := CreateConnection(*dbConfig)
	if err != nil {
		fmt.Println("Problemas ao conectar ao Banco de Dados!")
		return nil
	}
	fmt.Println("Conectado ao banco com sucesso!")
	return db
}

func CloseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Erro ao obter o objeto *sql.DB:", err)
		return
	}

	err = sqlDB.Close()
	if err != nil {
		fmt.Println("Erro ao fechar a conexão com o Banco de Dados:", err)
	}
	fmt.Println("Fechado conecção com o Banco de Dados!")
}
