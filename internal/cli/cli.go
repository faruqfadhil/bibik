package cli

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/briandowns/spinner"
	"github.com/faruqfadhil/bibik/handler/cli"
	"github.com/faruqfadhil/bibik/internal/cli/repository"
	"github.com/faruqfadhil/bibik/internal/entity"
)

type cliService struct {
	repo repository.CLIRepository
	spin *spinner.Spinner
}

func NewCLIService(repo repository.CLIRepository, spin *spinner.Spinner) cli.Service {
	return &cliService{repo: repo, spin: spin}
}

func (s *cliService) UpsertCommand(ctx context.Context, req *entity.Command) error {
	err := s.repo.Upsert(&repository.CLIModel{
		Key:   req.Key,
		Value: req.Value,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *cliService) GetByKey(ctx context.Context, key string) (*entity.Command, error) {
	resp, err := s.repo.FindByKey(key)
	if err != nil {
		return nil, err
	}

	return &entity.Command{
		Key:   resp.Key,
		Value: resp.Value,
	}, nil
}

func (s *cliService) SearchByKey(ctx context.Context, key string) ([]*entity.Command, error) {
	out := []*entity.Command{}
	datas, err := s.repo.SearchByKey(key)
	if err != nil {
		return nil, err
	}
	for _, d := range datas {
		out = append(out, &entity.Command{
			Key:   d.Key,
			Value: d.Value,
		})
	}
	return out, nil
}
func (s *cliService) Exec(ctx context.Context, key string) (string, error) {
	command, err := s.repo.FindByKey(key)
	if err != nil {
		return "", err
	}

	args := strings.Split(command.Value, " ")

	s.spin.Suffix = fmt.Sprintf(" Executing: %s ...", args)
	out, err := exec.Command(args[0], args[1:]...).Output()
	fmt.Println(" done")

	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (s *cliService) GetAll(ctx context.Context) ([]*entity.Command, error) {
	out := []*entity.Command{}
	datas, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	for _, d := range datas {
		out = append(out, &entity.Command{
			Key:   d.Key,
			Value: d.Value,
		})
	}
	return out, nil
}
