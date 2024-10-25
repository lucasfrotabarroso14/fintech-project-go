package account

import (
	"database/sql"
	"processamento-pagamento-go/internal/domain/entity/account_entity"
)

type AccountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (ar *AccountRepository) CreateAccount(account *account_entity.Account) error {
	query := "INSERT INTO accounts(id,user_id,balance) values (?,?,?)"
	_, err := ar.DB.Exec(query, account.Id, account.User_id, account.Balance)
	if err != nil {
		return err
	}
	return nil

}
