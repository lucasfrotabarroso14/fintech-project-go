package user_interface

import "processamento-pagamento-go/internal/domain/dto/user"

type UserUseCaseInterface interface {
	CreateUser(user *user.CreateUserDTO) error
}
