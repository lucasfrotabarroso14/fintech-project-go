package account_interface

type AccountUserCaseInterface interface {
	CreateAccount(accountId string) error
}
