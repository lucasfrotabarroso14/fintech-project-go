package transaction_entity

import "github.com/google/uuid"

type TransactionEntityInputDTO struct {
	From_account_id string
	To_account_id   string
	Amount          float64
}

type TransactionEntity struct {
	Id              string
	From_account_id string
	To_account_id   string
	Amount          float64
}

func CreateNewTransactionEntity(input TransactionEntityInputDTO) (*TransactionEntity, error) {
	return &TransactionEntity{
		Id:              uuid.New().String(),
		From_account_id: input.From_account_id,
		To_account_id:   input.To_account_id,
		Amount:          input.Amount,
	}, nil
}
