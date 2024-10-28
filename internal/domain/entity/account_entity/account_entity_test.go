package account_entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccount(t *testing.T) {

	accountTestEntity, err := NewAccount("123")
	assert.NoError(t, err)
	assert.Equal(t, "123", accountTestEntity.User_id)
	assert.Equal(t, 0.0, accountTestEntity.Balance)

}
