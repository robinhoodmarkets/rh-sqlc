// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package split

import (
	"context"

	"../models"
)

type Querier interface {
	GetStock(ctx context.Context) ([]models.Stock, error)
}

var _ Querier = (*Queries)(nil)