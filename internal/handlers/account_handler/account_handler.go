package account_handler

import (
	"encoding/json"
	"net/http"
	"processamento-pagamento-go/internal/domain/dto/account"
	"processamento-pagamento-go/internal/domain/interfaces/account_interface"
)

type AccountHandler struct {
	accountUseCase account_interface.AccountUserCaseInterface
}

func NewAccountHandler(auc account_interface.AccountUserCaseInterface) *AccountHandler {
	return &AccountHandler{
		accountUseCase: auc,
	}
}

func (ah *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var newAccount account.CreateAccountDTO

	if err := json.NewDecoder(r.Body).Decode(&newAccount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	if err := ah.accountUseCase.CreateAccount(newAccount.UserId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Conta criada com sucesso",
	})

}
