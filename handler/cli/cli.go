package cli

import (
	"context"

	"github.com/faruqfadhil/bibik/internal/entity"
)

type Service interface {
	UpsertCommand(ctx context.Context, req *entity.Command) error
	GetByKey(ctx context.Context, key string) (*entity.Command, error)
	SearchByKey(ctx context.Context, key string) ([]*entity.Command, error)
	Exec(ctx context.Context, key string) (string, error)
}
