package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"processamento-pagamento-go/internal/domain/adapters"
	"processamento-pagamento-go/internal/domain/usecase/transaction_usecase"
	"processamento-pagamento-go/internal/domain/usecase/user_usecase"
	"processamento-pagamento-go/internal/handlers/transaction_handler"
	"processamento-pagamento-go/internal/handlers/user_handler"
)

func main() {

	//err = dynamoDBRepo.CreateTransactionsTable()
	//if err != nil {
	//	log.Fatalf("Erro ao criar tabela transactions: %v", err)
	//}
	ctx := context.Background()
	adapter := adapters.New(ctx)

	defer adapter.DB.Close()

	mux := http.NewServeMux()

	transactionUseCase := transaction_usecase.NewTransactionUseCase(adapter.AccountRepo, adapter.TransactionRepo, adapter.DynamoDBRepo)
	transactionHandler := transaction_handler.NewTransactionHandler(transactionUseCase)

	userUseCase := user_usecase.NewUserUseCase(adapter.UserRepo, adapter.AccountRepo)
	userHandler := user_handler.NewUserHandler(userUseCase)

	mux.HandleFunc("/users", userHandler.CreateUser)

	mux.HandleFunc("/transaction", transactionHandler.CrateTransaction)

	if err := http.ListenAndServe(":8881", mux); err != nil {
		fmt.Printf("Error: %s", err)
	}

}
