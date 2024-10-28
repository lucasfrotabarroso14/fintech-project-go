package transaction_interface

import "processamento-pagamento-go/internal/domain/entity/transaction_entity"

type TransactionRepository interface {
	CreateNewTransaction(transaction *transaction_entity.TransactionEntity) error
}
