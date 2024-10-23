package interfaces

import (
	"processamento-pagamento-go/internal/domain/entity/user_entity"
)

type UserRepoInterface interface {
	CreateUser(user *user_entity.User) error
}
