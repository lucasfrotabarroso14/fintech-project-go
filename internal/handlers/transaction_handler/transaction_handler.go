package transaction_handler

import (
	"encoding/json"
	"net/http"
	"processamento-pagamento-go/internal/domain/dto/transaction_dto"
	"processamento-pagamento-go/internal/domain/interfaces/transaction_interface"
	"processamento-pagamento-go/pkg/responses"
)

type TransactionHandler struct {
	transactionUseCase transaction_interface.TransactionUsecase
}

func NewTransactionHandler(uc transaction_interface.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		transactionUseCase: uc,
	}
}

func (th *TransactionHandler) CrateTransaction(w http.ResponseWriter, r *http.Request) {
	var transactionDTO transaction_dto.TransactionInputDTO

	if err := json.NewDecoder(r.Body).Decode(&transactionDTO); err != nil {
		responses.Error(w, http.StatusBadRequest, "Error", err)
		return
	}
	if err := th.transactionUseCase.Execute(transactionDTO); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Error", err)
		return
	}
	responses.Success(w, http.StatusOK, "Success", nil)

}
