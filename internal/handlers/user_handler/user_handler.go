package user_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"processamento-pagamento-go/internal/domain/dto/user"
	"processamento-pagamento-go/internal/domain/interfaces"
)

type UserHandler struct {
	userUseCase interfaces.UserUseCaseInterface
}

func NewUserHandler(uc interfaces.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		userUseCase: uc,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("caiu no create!!!!!!")
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
