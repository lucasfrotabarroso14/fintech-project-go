package user_usecase

import (
	"fmt"
	"processamento-pagamento-go/internal/domain/dto/user"
	"processamento-pagamento-go/internal/domain/entity/account_entity"
	"processamento-pagamento-go/internal/domain/entity/user_entity"
	"processamento-pagamento-go/internal/domain/interfaces/account_interface"
	"processamento-pagamento-go/internal/domain/interfaces/user_interface"
)

type UserUseCase struct {
	userRepository    user_interface.UserRepoInterface
	accountRepository account_interface.AccountRepositoryInterface
}

func NewUserUseCase(userRepo user_interface.UserRepoInterface, accountRepo account_interface.AccountRepositoryInterface) *UserUseCase {
	return &UserUseCase{
		userRepository:    userRepo,
		accountRepository: accountRepo,
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

	accounEntity, err := account_entity.NewAccount(userEntity.Id)
	if err != nil {
		return err
	}

	if err = uc.accountRepository.CreateAccount(accounEntity); err != nil {
		return err

	}

	return nil

}
