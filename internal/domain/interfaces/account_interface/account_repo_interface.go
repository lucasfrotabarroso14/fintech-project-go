package account_interface

import (
	"database/sql"
	"processamento-pagamento-go/internal/domain/entity/account_entity"
)

type AccountRepositoryInterface interface {
	CreateAccount(account *account_entity.Account) error
	GetBalanceById(accountId string) (float64, error)
	GetById(accountId string) (*account_entity.Account, error)
	IncreaseBalance(tx *sql.Tx, accountId string, amount float64) error
	DecreaseBalance(tx *sql.Tx, accountId string, amount float64) error
	BeginTransaction() (*sql.Tx, error)
}
