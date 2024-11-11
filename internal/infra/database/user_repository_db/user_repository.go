package user_repository_db

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"processamento-pagamento-go/internal/domain/entity/user_entity"
	"processamento-pagamento-go/pkg/logger"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) CreateUser(user *user_entity.User) error {
	query := "INSERT INTO users(id, name, email, password_hash, cpf) VALUES (?, ?, ?, ?, ?)"
	_, err := ur.DB.Exec(query, user.Id, user.Name, user.Email, user.Password_hash, user.Cpf)
	if err != nil {
		logger.Log.Error("Failed to create user in MySQL",
			zap.String("user_id", user.Id),
			zap.String("name", user.Name),
			zap.String("email", user.Email),
			zap.String("cpf", user.Cpf),
			zap.Error(err),
		)
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
