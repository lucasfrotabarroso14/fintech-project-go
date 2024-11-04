package transaction_interface

import (
	"processamento-pagamento-go/internal/domain/entity/transaction_entity"
)

type TransactionDynamoDBRepoInterface interface {
	SaveTransaction(transaction *transaction_entity.TransactionEntity) error
	//CreateTransactionsTable() error
}
