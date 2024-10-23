package user

import (
	"fmt"
	"processamento-pagamento-go/internal/domain/dto/user"
	"processamento-pagamento-go/internal/domain/entity/user_entity"
	"processamento-pagamento-go/internal/domain/interfaces"
)

type UserUseCase struct {
	userRepository interfaces.UserRepoInterface
}

func NewUserUseCase(repo interfaces.UserRepoInterface) *UserUseCase {
	return &UserUseCase{
		userRepository: repo,
	}
}

func (uc *UserUseCase) CreateUser(user *user.CreateUserDTO) error {
	userEntity, err := user_entity.NewUser(user)
	if err != nil {
		fmt.Println("Error: create user entity -> ", err)
		return err
	}

	if err := uc.userRepository.CreateUser(userEntity); err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
