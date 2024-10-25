package account_entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccount(t *testing.T) {

	acountTestEntity, err := NewAccount("123")
	assert.NoError(t, err)
	assert.Equal(t, "123", acountTestEntity.User_id)
	assert.Equal(t, 0.0, acountTestEntity.Balance)

}
