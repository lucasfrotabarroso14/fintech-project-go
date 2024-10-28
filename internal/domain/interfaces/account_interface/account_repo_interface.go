package account_interface

import "processamento-pagamento-go/internal/domain/entity/account_entity"

type AccountRepositoryInterface interface {
	CreateAccount(account *account_entity.Account) error
	GetBalanceById(accountId string) (float64, error)
	GetById(accountId string) (*account_entity.Account, error)
	IncreaseBalance(accountId string, amount float64) error
	DecreaseBalance(accountId string, amount float64) error
}
