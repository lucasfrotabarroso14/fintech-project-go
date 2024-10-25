package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	account_usecase "processamento-pagamento-go/internal/domain/usecase/account_usecase"
	"processamento-pagamento-go/internal/domain/usecase/user_usecase"
	"processamento-pagamento-go/internal/handlers/account_handler"
	"processamento-pagamento-go/internal/handlers/user_handler"
	"processamento-pagamento-go/internal/infra/database"
	"processamento-pagamento-go/internal/infra/database/account"
	user_repository "processamento-pagamento-go/internal/infra/database/user"
)

func main() {

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

	accountRepo := account.NewAccountRepository(db)
	accountUserCase := account_usecase.NewAccountUserCase(accountRepo)
	accountHandler := account_handler.NewAccountHandler(accountUserCase)

	userRepo := user_repository.NewUserRepository(db)
	userUseCase := user_usecase.NewUserUseCase(userRepo)
	userHandler := user_handler.NewUserHandler(userUseCase)

	mux.HandleFunc("/users", userHandler.CreateUser)
	mux.HandleFunc("/account", accountHandler.CreateAccount)

	if err = http.ListenAndServe(":8881", mux); err != nil {
		fmt.Printf("Error: %s", err)
	}

}
