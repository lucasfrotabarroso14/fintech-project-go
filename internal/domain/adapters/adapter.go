package adapters

import (
	"context"
	"database/sql"
	"log"
	"processamento-pagamento-go/internal/domain/interfaces/account_interface"
	"processamento-pagamento-go/internal/domain/interfaces/transaction_interface"
	"processamento-pagamento-go/internal/domain/interfaces/user_interface"
	"processamento-pagamento-go/internal/infra/database"
	"processamento-pagamento-go/internal/infra/database/account_repository_db"
	"processamento-pagamento-go/internal/infra/database/dynamodb_repository"
	"processamento-pagamento-go/internal/infra/database/transaction_repository_db"
	user_repository "processamento-pagamento-go/internal/infra/database/user_repository_db"
)

type Adapters struct {
	TransactionRepo transaction_interface.TransactionRepository
	AccountRepo     account_interface.AccountRepositoryInterface
	UserRepo        user_interface.UserRepoInterface
	DynamoDBRepo    transaction_interface.TransactionDynamoDBRepoInterface
	DB              *sql.DB
}

func New(ctx context.Context) *Adapters {
	adapters := &Adapters{}
	dsn := database.GenerateDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error database connection:", err)
	}
	err = db.Ping()
	adapters.DB = db

	if err != nil {
		log.Fatal("error database ping ", err)

	}

	dynamoDBRepo, err := dynamodb_repository.NewDynamoDBRepository()
	if err != nil {
		log.Fatal(err)
	}
	adapters.DynamoDBRepo = dynamoDBRepo

	transactionRepo := transaction_repository_db.NewTransactionRepository(db)
	adapters.TransactionRepo = transactionRepo

	accountRepo := account_repository_db.NewAccountRepository(db)
	adapters.AccountRepo = accountRepo

	userRepo := user_repository.NewUserRepository(db)
	adapters.UserRepo = userRepo

	return adapters
}
