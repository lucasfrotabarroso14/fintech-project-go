package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func GenerateDSN() string {
	//const envFilePath = "C:/Users/lfrota/GolandProjects/processamento-pagamento-go/.env"
	// Carregar o arquivo .env
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	// Constrói o caminho absoluto para o arquivo .env subindo 4 níveis
	envFilePath := filepath.Join(wd, "../../../../.env") // Ajuste o número de ".." conforme necessário

	// Carregar o arquivo .env
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatalf("Error to load .env file from %s -> %v", envFilePath, err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
}