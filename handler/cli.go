package handler

import (
	"context"

	"github.com/faruqfadhil/bibik/handler/cli"
	"github.com/faruqfadhil/bibik/internal/entity"
)

type CLIHandler struct {
	service cli.Service
}

func NewCLIHandler(service cli.Service) *CLIHandler {
	return &CLIHandler{
		service: service,
	}
}

func (h *CLIHandler) UpsertCommand(ctx context.Context, req *entity.Command) error {
	return h.service.UpsertCommand(ctx, req)
}

func (h *CLIHandler) GetByKey(ctx context.Context, key string) (*entity.Command, error) {
	return h.service.GetByKey(ctx, key)
}

func (h *CLIHandler) SearchByKey(ctx context.Context, key string) ([]*entity.Command, error) {
	return h.service.SearchByKey(ctx, key)
}

func (h *CLIHandler) Exec(ctx context.Context, key string, useDirr bool) (string, error) {
	return h.service.Exec(ctx, key, useDirr)
}

func (h *CLIHandler) GetAll(ctx context.Context) ([]*entity.Command, error) {
	return h.service.GetAll(ctx)
}
