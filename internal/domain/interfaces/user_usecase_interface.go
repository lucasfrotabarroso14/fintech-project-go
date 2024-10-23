package interfaces

import "processamento-pagamento-go/internal/domain/dto/user"

type UserUseCaseInterface interface {
	CreateUser(user *user.CreateUserDTO) error
}
