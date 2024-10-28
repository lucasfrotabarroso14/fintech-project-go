package user_repository_db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"processamento-pagamento-go/internal/domain/entity/user_entity"
	"processamento-pagamento-go/internal/infra/database"
	"testing"
)

func TestUserRepository_CreateUse(t *testing.T) {
	dsn := database.GenerateDSN()
	db, err := sql.Open("mysql", dsn)
	assert.NoError(t, err, "Failed to connect db")
	newUUID := uuid.New().String()

	userRepo := NewUserRepository(db)

	testUser := &user_entity.User{
		Id:            newUUID,
		Name:          "John Doe",
		Email:         "john@example.com",
		Password_hash: "hashed_password",
		Cpf:           "12345678901",
	}

	err = userRepo.CreateUser(testUser)
	assert.NoError(t, err, "Error creating user_usecase")

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE id =?", testUser.Id).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	_, err = db.Exec("DELETE FROM users WHERE ID = ?", testUser.Id)
	assert.NoError(t, err)

}
