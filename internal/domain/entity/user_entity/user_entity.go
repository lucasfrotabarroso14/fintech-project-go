package user_entity

import (
	"fmt"
	"github.com/google/uuid"
	"processamento-pagamento-go/internal/domain/dto/user"
	"processamento-pagamento-go/pkg/auth"
)

type User struct {
	Id            string
	Name          string
	Email         string
	Password_hash string
	Cpf           string
}

// criar depois funcao de validacao dos campos

func NewUser(user *user.CreateUserDTO) (*User, error) {

	hashPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return nil, err

	}

	newUser := &User{
		Id:            uuid.New().String(),
		Name:          user.Name,
		Email:         user.Email,
		Password_hash: hashPassword,
		Cpf:           user.Cpf,
	}

	return newUser, nil

}
