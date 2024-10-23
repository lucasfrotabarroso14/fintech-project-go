package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"processamento-pagamento-go/internal/domain/interfaces"
	"processamento-pagamento-go/internal/domain/usecase/user"
	"processamento-pagamento-go/internal/handlers/user_handler"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working directory:", err)
	}
	fmt.Println("Current working directory:", wd)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error to load .env file -> ", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error database connection:", err)
	}

	defer db.Close()
	err = db.Ping()

	if err != nil {
		log.Fatal("error database ping ", err)

	}
	mux := http.NewServeMux()

	var userRepo interfaces.UserRepoInterface
	userUseCase := user.NewUserUseCase(userRepo)
	userHandler := user_handler.NewUserHandler(userUseCase)
	mux.HandleFunc("/users", userHandler.CreateUser)

	if err := http.ListenAndServe(":8881", mux); err != nil {
		fmt.Printf("Error: %s", err)
	}

}
