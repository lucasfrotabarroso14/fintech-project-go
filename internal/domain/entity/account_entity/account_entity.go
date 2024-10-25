package account_entity

import "github.com/google/uuid"

type Account struct {
	Id      string
	User_id string
	Balance float64
}

func NewAccount(newUserId string) (*Account, error) {
	return &Account{
		Id:      uuid.New().String(),
		User_id: newUserId,
		Balance: 0.0,
	}, nil

}
