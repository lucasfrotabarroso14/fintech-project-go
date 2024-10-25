package user

import (
	"database/sql"
	"fmt"
	"processamento-pagamento-go/internal/domain/entity/user_entity"
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
	query := "INSERT INTO users(id,name,email, password_hash,cpf ) VALUES (?,?,?,?,?)"
	_, err := ur.DB.Exec(query, user.Id, user.Name, user.Email, user.Password_hash, user.Cpf)
	if err != nil {
		return fmt.Errorf("failed to create user_usecase: %w", err)
	}
	return nil
}
