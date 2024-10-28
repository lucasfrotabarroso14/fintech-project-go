package user_handler

import (
	"encoding/json"
	"net/http"
	"processamento-pagamento-go/internal/domain/dto/user"
	"processamento-pagamento-go/internal/domain/interfaces/user_interface"
	"processamento-pagamento-go/pkg/responses"
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
		responses.Error(w, http.StatusBadRequest, "Error", err)
		return
	}
	if err := uh.userUseCase.CreateUser(&userDTO); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Error", err)
		return
	}

	responses.Success(w, http.StatusCreated, "Success", nil)

}
