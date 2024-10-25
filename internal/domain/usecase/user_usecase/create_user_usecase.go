package user_usecase

import (
	"fmt"
	"processamento-pagamento-go/internal/domain/dto/user"
	"processamento-pagamento-go/internal/domain/entity/user_entity"
	"processamento-pagamento-go/internal/domain/interfaces/user_interface"
)

type UserUseCase struct {
	userRepository user_interface.UserRepoInterface
}

func NewUserUseCase(repo user_interface.UserRepoInterface) *UserUseCase {
	return &UserUseCase{
		userRepository: repo,
	}
}

func (uc *UserUseCase) CreateUser(user *user.CreateUserDTO) error {
	userEntity, err := user_entity.NewUser(user)
	if err != nil {
		fmt.Println("Error: create user_usecase entity -> ", err)
		return err
	}

	if err = uc.userRepository.CreateUser(userEntity); err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
