package transaction_usecase

import (
	"errors"
	"processamento-pagamento-go/internal/domain/dto/transaction_dto"
	"processamento-pagamento-go/internal/domain/entity/transaction_entity"
	"processamento-pagamento-go/internal/domain/interfaces/account_interface"
	"processamento-pagamento-go/internal/domain/interfaces/transaction_interface"
)

// vou realizar a transferencia pelo id
//verificar se a conta de quem esta enviando tem saldo
// realizar a transferencia da conta x para a y

type TransactionUseCase struct {
	accountRepository       account_interface.AccountRepositoryInterface
	transactionRepository   transaction_interface.TransactionRepository
	transactionDynamoDBRepo transaction_interface.TransactionDynamoDBRepoInterface
}

func NewTransactionUseCase(accountRepo account_interface.AccountRepositoryInterface, transactionRepo transaction_interface.TransactionRepository, transactionDynamoDBRepo transaction_interface.TransactionDynamoDBRepoInterface) *TransactionUseCase {
	return &TransactionUseCase{
		accountRepository:       accountRepo,
		transactionRepository:   transactionRepo,
		transactionDynamoDBRepo: transactionDynamoDBRepo,
	}
}

func (uc *TransactionUseCase) Execute(input transaction_dto.TransactionInputDTO) error {

	transactionEntityDTO := transaction_entity.TransactionEntityInputDTO{
		From_account_id: input.From_account_id,
		To_account_id:   input.To_account_id,
		Amount:          input.Amount,
	}
	// receber os dados
	transactionEntity, err := transaction_entity.CreateNewTransactionEntity(transactionEntityDTO)
	if err != nil {
		return err
	}

	// verificar se a conta do from existe e se a conta do to existe\
	_, err = uc.accountRepository.GetById(input.From_account_id)
	if err != nil {
		return err
	}

	_, err = uc.accountRepository.GetById(input.To_account_id)
	if err != nil {
		return err
	}

	// verificar se a conta do from tem valor suficiente para fazer a transacao
	accountBalance, err := uc.accountRepository.GetBalanceById(input.From_account_id)
	if err != nil {
		return err
	}
	if accountBalance < input.Amount {
		return errors.New("Insuficient Balance")
	}

	//subtrair da conta do from

	if err = uc.accountRepository.IncreaseBalance(input.From_account_id, input.Amount); err != nil {
		return err
	}

	if err = uc.transactionRepository.CreateNewTransaction(transactionEntity); err != nil {
		return err
	}

	if err = uc.accountRepository.DecreaseBalance(input.To_account_id, input.Amount); err != nil {
		return err
	}

	if err = uc.transactionDynamoDBRepo.SaveTransaction(transactionEntity); err != nil {
		return err
	}

	return nil

	// somar da conta do to

}
