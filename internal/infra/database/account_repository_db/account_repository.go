package account_repository_db

import (
	"database/sql"
	"go.uber.org/zap"
	"processamento-pagamento-go/internal/domain/entity/account_entity"
	"processamento-pagamento-go/pkg/logger"
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
	query := "INSERT INTO accounts(id, user_id, balance) VALUES (?, ?, ?)"
	_, err := ar.DB.Exec(query, account.Id, account.User_id, account.Balance)
	if err != nil {
		logger.Log.Error("Failed to create account in MySQL",
			zap.String("account_id", account.Id),
			zap.String("user_id", account.User_id),
			zap.Float64("balance", account.Balance),
			zap.Error(err),
		)
		return err
	}
	return nil
}

func (ar *AccountRepository) GetBalanceById(accountId string) (float64, error) {
	var balance float64
	query := "SELECT balance FROM accounts WHERE id = ?"
	err := ar.DB.QueryRow(query, accountId).Scan(&balance)
	if err != nil {
		logger.Log.Error("Failed to retrieve balance from MySQL",
			zap.String("account_id", accountId),
			zap.Error(err),
		)
		return 0, err
	}
	return balance, nil
}

func (ar *AccountRepository) IncreaseBalance(accountId string, amount float64) error {
	query := "UPDATE accounts SET balance = balance + ? WHERE id = ?"
	_, err := ar.DB.Exec(query, amount, accountId)
	if err != nil {
		logger.Log.Error("Failed to increase balance in MySQL",
			zap.String("account_id", accountId),
			zap.Float64("amount", amount),
			zap.Error(err),
		)
		return err
	}
	return nil
}

func (ar *AccountRepository) DecreaseBalance(accountId string, amount float64) error {
	query := "UPDATE accounts SET balance = balance - ? WHERE id = ?"
	_, err := ar.DB.Exec(query, amount, accountId)
	if err != nil {
		logger.Log.Error("Failed to decrease balance in MySQL",
			zap.String("account_id", accountId),
			zap.Float64("amount", amount),
			zap.Error(err),
		)
		return err
	}
	return nil
}

func (ar *AccountRepository) GetById(accountId string) (*account_entity.Account, error) {
	var account account_entity.Account
	query := "SELECT id, user_id, balance FROM accounts WHERE id = ?"
	err := ar.DB.QueryRow(query, accountId).Scan(&account.Id, &account.User_id, &account.Balance)
	if err != nil {
		logger.Log.Error("Failed to retrieve account from MySQL",
			zap.String("account_id", accountId),
			zap.Error(err),
		)
		return nil, err
	}
	return &account, nil
}
