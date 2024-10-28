package transaction_repository_db

import (
	"database/sql"
	"processamento-pagamento-go/internal/domain/entity/transaction_entity"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (tr *TransactionRepository) CreateNewTransaction(transaction *transaction_entity.TransactionEntity) error {
	query := "INSERT INTO transactions(id,from_account_id,to_account_id,amount) values (?,?,?,?)"
	_, err := tr.DB.Exec(query, transaction.Id, transaction.From_account_id, transaction.To_account_id, transaction.Amount)
	if err != nil {
		return err
	}
	return nil

}
