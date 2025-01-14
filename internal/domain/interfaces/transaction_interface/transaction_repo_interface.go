package transaction_interface

import (
	"database/sql"
	"processamento-pagamento-go/internal/domain/entity/transaction_entity"
)

type TransactionRepository interface {
	CreateNewTransaction(tx *sql.Tx, transaction *transaction_entity.TransactionEntity) error
}
