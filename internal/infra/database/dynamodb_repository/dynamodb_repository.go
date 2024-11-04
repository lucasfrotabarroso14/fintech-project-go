package dynamodb_repository

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"processamento-pagamento-go/internal/domain/entity/transaction_entity"
)

type DynamoDBRepository struct {
	Client *dynamodb.Client
}

func NewDynamoDBRepository() (*DynamoDBRepository, error) {
	const endpoint = "http://localhost:8000" // URL do DynamoDB Local
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
		return nil, fmt.Errorf("error loading configuration: %w", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDBRepository{Client: client}, nil
}

func (dr *DynamoDBRepository) SaveTransaction(transaction *transaction_entity.TransactionEntity) error {
	// Converte a entidade da transação para um mapa de atributos DynamoDB
	item, err := attributevalue.MarshalMap(transaction)
	if err != nil {
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
		return fmt.Errorf("failed to save transaction in DynamoDB: %w", err)
	}

	return nil
}

//func (dr *DynamoDBRepository) CreateTransactionsTable() error {
//	// Define a estrutura da tabela
//	input := &dynamodb.CreateTableInput{
//		TableName: aws.String("transactions"),
//		AttributeDefinitions: []types.AttributeDefinition{
//			{
//				AttributeName: aws.String("id"),
//				AttributeType: types.ScalarAttributeTypeS, // Tipo STRING para o id
//			},
//		},
//		KeySchema: []types.KeySchemaElement{
//			{
//				AttributeName: aws.String("id"),
//				KeyType:       types.KeyTypeHash, // id é a chave HASH (partição)
//			},
//		},
//		ProvisionedThroughput: &types.ProvisionedThroughput{
//			ReadCapacityUnits:  aws.Int64(5),
//			WriteCapacityUnits: aws.Int64(5),
//		},
//	}
//
//	// Tenta criar a tabela
//	_, err := dr.Client.CreateTable(context.TODO(), input)
//	if err != nil {
//		log.Printf("Falha ao criar tabela: %v", err)
//		return err
//	}
//
//	log.Println("Tabela 'transactions' criada com sucesso")
//	return nil
//}
