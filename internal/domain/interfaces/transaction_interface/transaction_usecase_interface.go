package transaction_interface

import (
	"processamento-pagamento-go/internal/domain/dto/transaction_dto"
)

type TransactionUsecase interface {
	Execute(input transaction_dto.TransactionInputDTO) error
}
