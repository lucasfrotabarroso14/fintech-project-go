package transaction_dto

type TransactionInputDTO struct {
	From_account_id string  `json:"from_account_id"`
	To_account_id   string  `json:"to_account_id"`
	Amount          float64 `json:"amount"`
}
