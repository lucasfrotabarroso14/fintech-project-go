package user_entity

import (
	"github.com/stretchr/testify/assert"
	"processamento-pagamento-go/internal/domain/dto/user"
	"testing"
)

func TestNewUser(t *testing.T) {

	userDTO := user.CreateUserDTO{
		Name:     "Lucas",
		Email:    "lucas@gmail.com",
		Password: "1234567",
		Cpf:      "02122859402",
	}

	u, err := NewUser(&userDTO)
	assert.Nil(t, err)
	assert.Equal(t, "Lucas", u.Name)
	assert.Equal(t, "lucas@gmail.com", u.Email)
	assert.Equal(t, "02122859402", u.Cpf)

}
