package repository

import (
	"context"

	"github.com/MatsuoTakuro/ddd-article/domain"
)

type Task interface {
	Save(context.Context, domain.Task) error
	GetByID(ctx context.Context, taskID string) (domain.Task, error)
}
