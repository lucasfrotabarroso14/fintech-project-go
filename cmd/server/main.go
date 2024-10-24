package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"processamento-pagamento-go/internal/domain/usecase/user"
	"processamento-pagamento-go/internal/handlers/user_handler"
	"processamento-pagamento-go/internal/infra/database"
	user_repository "processamento-pagamento-go/internal/infra/database/user"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working directory:", err)
	}
	fmt.Println("Current working directory:", wd)

	dsn := database.GenerateDSN()

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

	userRepo := user_repository.NewUserRepository(db)
	userUseCase := user.NewUserUseCase(userRepo)
	userHandler := user_handler.NewUserHandler(userUseCase)
	mux.HandleFunc("/users", userHandler.CreateUser)

	if err := http.ListenAndServe(":8881", mux); err != nil {
		fmt.Printf("Error: %s", err)
	}

}
