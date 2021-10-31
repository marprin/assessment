package storage

import (
	"context"

	"github.com/marprin/assessment/fetchapp/internal/domain/storage/entity"
)

type (
	Currency interface {
		FetchCurrencyConverter(ctx context.Context, q string) (*float64, error)
	}

	Service interface {
		StorageList(ctx context.Context) ([]entity.StorageListResponse, error)
	}

	Gateway interface {
		FetchStorageList(ctx context.Context) ([]entity.GatewayListResponse, error)
	}
)
