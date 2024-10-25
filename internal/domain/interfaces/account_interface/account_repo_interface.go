package account_interface

import "processamento-pagamento-go/internal/domain/entity/account_entity"

type AccountRepositoryInterface interface {
	CreateAccount(account *account_entity.Account) error
}
