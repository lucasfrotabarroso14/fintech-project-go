package transaction_repository_db

import (
	"database/sql"
	"go.uber.org/zap"
	"processamento-pagamento-go/internal/domain/entity/transaction_entity"
	"processamento-pagamento-go/pkg/logger"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (tr *TransactionRepository) CreateNewTransaction(tx *sql.Tx, transaction *transaction_entity.TransactionEntity) error {
	query := "INSERT INTO transactions(id, from_account_id, to_account_id, amount) VALUES (?, ?, ?, ?)"
	_, err := tx.Exec(query, transaction.Id, transaction.From_account_id, transaction.To_account_id, transaction.Amount)
	if err != nil {
		logger.Log.Error("Failed to create new transaction in MySQL",
			zap.String("transaction_id", transaction.Id),
			zap.String("from_account_id", transaction.From_account_id),
			zap.String("to_account_id", transaction.To_account_id),
			zap.Float64("amount", transaction.Amount),
			zap.Error(err),
		)
		return err
	}
	return nil
}
