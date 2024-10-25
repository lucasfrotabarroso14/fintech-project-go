package account_usecase

import (
	"processamento-pagamento-go/internal/domain/entity/account_entity"
	"processamento-pagamento-go/internal/domain/interfaces/account_interface"
)

type AccountUserCase struct {
	AccountRepo account_interface.AccountRepositoryInterface
}

func NewAccountUserCase(accountRepository account_interface.AccountRepositoryInterface) *AccountUserCase {
	return &AccountUserCase{
		AccountRepo: accountRepository,
	}
}

func (au *AccountUserCase) CreateAccount(accountId string) error {
	newAccount := account_entity.Account{
		User_id: accountId,
	}
	if err := au.AccountRepo.CreateAccount(&newAccount); err != nil {
		return err

	}
	return nil

}
