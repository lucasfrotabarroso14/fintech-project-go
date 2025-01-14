package dynamodb_repository

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"go.uber.org/zap"
	"processamento-pagamento-go/internal/domain/entity/transaction_entity"
	"processamento-pagamento-go/pkg/logger"
)

type DynamoDBRepository struct {
	Client *dynamodb.Client
}

func NewDynamoDBRepository() (*DynamoDBRepository, error) {
	const endpoint = "http://localhost:8010" // URL do DynamoDB Local
	const region = "us-west-2"               // Região configurada

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			// Aplica o endpoint apenas para o serviço DynamoDB
			if service == dynamodb.ServiceID {
				return aws.Endpoint{URL: endpoint}, nil
			}
			return aws.Endpoint{}, fmt.Errorf("unsupported endpoint for service: %s", service)
		})),
	)
	if err != nil {
		logger.Log.Error("Error loading DynamoDB configuration",
			zap.String("endpoint", endpoint),
			zap.String("region", region),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error loading configuration: %w", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDBRepository{Client: client}, nil
}

func (dr *DynamoDBRepository) SaveTransaction(transaction *transaction_entity.TransactionEntity) error {
	// Converte a entidade da transação para um mapa de atributos DynamoDB
	item, err := attributevalue.MarshalMap(transaction)
	if err != nil {
		logger.Log.Error("Error marshaling transaction entity",
			zap.String("transaction_id", transaction.Id),
			zap.Error(err),
		)
		return fmt.Errorf("error marshaling transaction: %w", err)
	}

	// Prepara a entrada do item para inserir na tabela 'transactions'
	input := &dynamodb.PutItemInput{
		TableName: aws.String("transactions"),
		Item:      item,
	}

	// Executa a operação PutItem e trata o erro, se ocorrer
	_, err = dr.Client.PutItem(context.TODO(), input)
	if err != nil {
		logger.Log.Error("Failed to save transaction in DynamoDB",
			zap.String("transaction_id", transaction.Id),
			zap.Error(err),
		)
		return fmt.Errorf("failed to save transaction in DynamoDB: %w", err)
	}

	return nil
}
