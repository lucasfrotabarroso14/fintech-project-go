package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"processamento-pagamento-go/internal/domain/usecase/transaction_usecase"
	"processamento-pagamento-go/internal/infra/database/dynamodb_repository"

	"processamento-pagamento-go/internal/domain/usecase/user_usecase"
	"processamento-pagamento-go/internal/handlers/transaction_handler"
	"processamento-pagamento-go/internal/handlers/user_handler"
	"processamento-pagamento-go/internal/infra/database"
	"processamento-pagamento-go/internal/infra/database/account_repository_db"
	"processamento-pagamento-go/internal/infra/database/transaction_repository_db"
	user_repository "processamento-pagamento-go/internal/infra/database/user_repository_db"
)

func main() {

	dsn := database.GenerateDSN()

	dynamoDBRepo, err := dynamodb_repository.NewDynamoDBRepository()
	if err != nil {
		log.Fatal(err)
	}
	//err = dynamoDBRepo.CreateTransactionsTable()
	//if err != nil {
	//	log.Fatalf("Erro ao criar tabela transactions: %v", err)
	//}
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

	transactionRepo := transaction_repository_db.NewTransactionRepository(db)
	accountRepo := account_repository_db.NewAccountRepository(db)
	userRepo := user_repository.NewUserRepository(db)

	transactionUseCase := transaction_usecase.NewTransactionUseCase(accountRepo, transactionRepo, dynamoDBRepo)
	transactionHandler := transaction_handler.NewTransactionHandler(transactionUseCase)

	userUseCase := user_usecase.NewUserUseCase(userRepo, accountRepo)
	userHandler := user_handler.NewUserHandler(userUseCase)

	mux.HandleFunc("/users", userHandler.CreateUser)

	mux.HandleFunc("/transaction", transactionHandler.CrateTransaction)

	if err = http.ListenAndServe(":8881", mux); err != nil {
		fmt.Printf("Error: %s", err)
	}

}
