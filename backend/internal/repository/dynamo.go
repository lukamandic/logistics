package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ParcelSize []float64

type TableItem struct {
	ID          string     `json:"id"`
	ParcelSizes ParcelSize `json:"parcel_sizes"`
}

type DynamoRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoClient(ctx context.Context, tableName string) (*DynamoRepository, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %w", err)
	}

	client := dynamodb.NewFromConfig(cfg)

	return &DynamoRepository{
		client:    client,
		tableName: tableName,
	}, nil
}

func (r *DynamoRepository) GetAllItems(ctx context.Context) ([]TableItem, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	}

	result, err := r.client.Scan(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan table: %w", err)
	}

	fmt.Println("result", result.Items)

	var items []TableItem
	for _, item := range result.Items {
		tableItem, err := unmarshalTableItem(item)
		fmt.Println("tableitem", tableItem)
		if err != nil {
			fmt.Println("error", err)
			return nil, fmt.Errorf("failed to unmarshal item: %w", err)
		}
		items = append(items, tableItem)
	}

	return items, nil
}

func (r *DynamoRepository) PutItem(ctx context.Context, item TableItem) error {
	av, err := marshalTableItem(item)
	if err != nil {
		return fmt.Errorf("failed to marshal item: %w", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      av,
	}

	_, err = r.client.PutItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to put item: %w", err)
	}

	return nil
}

func (r *DynamoRepository) DeleteItem(ctx context.Context, id string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	}

	_, err := r.client.DeleteItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}

	return nil
}

func (r *DynamoRepository) UpsertItem(ctx context.Context, item TableItem) error {
	av, err := marshalTableItem(item)
	if err != nil {
		return fmt.Errorf("failed to marshal item: %w", err)
	}

	putInput := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      av,
	}

	_, err = r.client.PutItem(ctx, putInput)
	if err != nil {
		return fmt.Errorf("failed to upsert item: %w", err)
	}

	return nil
}

func unmarshalTableItem(av map[string]types.AttributeValue) (TableItem, error) {
	var item TableItem

	if id, ok := av["id"].(*types.AttributeValueMemberS); ok {
		item.ID = id.Value
	} else {
		return TableItem{}, fmt.Errorf("invalid id attribute")
	}

	if sizes, ok := av["parcel_sizes"].(*types.AttributeValueMemberS); ok {
		if err := json.Unmarshal([]byte(sizes.Value), &item.ParcelSizes); err != nil {
			return TableItem{}, fmt.Errorf("failed to parse parcel sizes JSON: %w", err)
		}
	} else {
		return TableItem{}, fmt.Errorf("invalid parcel_sizes attribute")
	}

	return item, nil
}

func marshalTableItem(item TableItem) (map[string]types.AttributeValue, error) {
	parcelSizesJSON, err := json.Marshal(item.ParcelSizes)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal parcel sizes: %w", err)
	}

	return map[string]types.AttributeValue{
		"id":           &types.AttributeValueMemberS{Value: item.ID},
		"parcel_sizes": &types.AttributeValueMemberS{Value: string(parcelSizesJSON)},
	}, nil
} 