// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package split

import (
	"context"

	"../models"
)

type Querier interface {
	GetWarehouse(ctx context.Context) ([]models.Warehouse, error)
}

var _ Querier = (*Queries)(nil)
