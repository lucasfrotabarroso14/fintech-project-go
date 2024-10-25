package user_handler

import (
	"encoding/json"
	"net/http"
	"processamento-pagamento-go/internal/domain/dto/user"
	"processamento-pagamento-go/internal/domain/interfaces/user_interface"
)

type UserHandler struct {
	userUseCase user_interface.UserUseCaseInterface
}

func NewUserHandler(uc user_interface.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		userUseCase: uc,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var userDTO user.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := uh.userUseCase.CreateUser(&userDTO); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "usuario criado com sucesso",
	})

}
