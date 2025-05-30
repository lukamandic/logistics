package service

import (
	"context"
	"fmt"

	"github.com/lukamandic/logistics/backend/internal/repository"
	"github.com/lukamandic/logistics/backend/internal/utils"
)

type ParcelSize = repository.ParcelSize

type TableItem = repository.TableItem

type ParcelService struct {
	repo *repository.DynamoRepository
}

func NewParcelService(repo *repository.DynamoRepository) *ParcelService {
	return &ParcelService{repo: repo}
}

func (s *ParcelService) GetAllItems(ctx context.Context) ([]TableItem, error) {
	return s.repo.GetAllItems(ctx)
}

func (s *ParcelService) UpdateParcelSizes(ctx context.Context, id string, sizes ParcelSize) error {
	newItem := TableItem{
		ID:          id,
		ParcelSizes: sizes,
	}

	if err := s.repo.UpsertItem(ctx, newItem); err != nil {
		return fmt.Errorf("failed to update parcel sizes: %w", err)
	}

	return nil
}

func (s *ParcelService) CalculateDistribution(ctx context.Context, amount float64) (map[float64]int, error) {
	items, err := s.repo.GetAllItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch items: %w", err)
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("no parcel sizes found in the table")
	}

	if len(items) > 1 {
		return nil, fmt.Errorf("multiple items found in the table, expected only one")
	}

	parcelSizes := items[0].ParcelSizes
	packSizes := make([]int, len(parcelSizes))
	for i, size := range parcelSizes {
		packSizes[i] = int(size)
	}

	orderedAmount := int(amount)
	if orderedAmount <= 0 {
		return nil, fmt.Errorf("ordered amount must be greater than zero")
	}

	distribution := utils.PackageDistribution(packSizes, orderedAmount)
	if distribution == nil {
		return nil, fmt.Errorf("no valid distribution found")
	}

	result := make(map[float64]int)
	for _, size := range parcelSizes {
		result[size] = 0
	}

	for size, count := range distribution {
		result[float64(size)] = count
	}

	return result, nil
} 