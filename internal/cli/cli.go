package cli

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/briandowns/spinner"
	"github.com/faruqfadhil/bibik/handler/cli"
	"github.com/faruqfadhil/bibik/internal/cli/repository"
	"github.com/faruqfadhil/bibik/internal/entity"
	errLib "github.com/faruqfadhil/bibik/pkg/error"
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

	if req.Options != nil && req.Options.Dir != "" {
		// save dir as a new data.
		err := s.repo.Upsert(&repository.CLIModel{
			Key:   req.SetDirKey(),
			Value: req.Options.Dir,
		})
		if err != nil {
			go s.repo.DeleteByKey(req.Key)
			return err
		}
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
func (s *cliService) Exec(ctx context.Context, key string, useDirr bool) (string, error) {
	command, err := s.repo.FindByKey(key)
	if err != nil {
		return "", err
	}

	var dirrCommandToBeExec string
	if useDirr {
		c := entity.Command{
			Key: key,
		}
		dir, err := s.repo.FindByKey(c.SetDirKey())
		if err != nil && !errors.Is(err, errLib.ErrKeyNotFound) {
			return "", err
		}
		if dir == nil || (dir != nil && dir.Value == "") {
			return "", errLib.ErrExecDirrNotFound
		}
		dirrCommandToBeExec = dir.Value
	}
	args := strings.Split(command.Value, " ")
	cmd := exec.Command(args[0], args[1:]...)
	if dirrCommandToBeExec != "" {
		cmd.Dir = dirrCommandToBeExec
	}
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err = cmd.Run()
	if err != nil {
		return "", err
	}
	outStr, errStr := stdoutBuf.String(), stderrBuf.String()

	if errStr != "" {
		return "", errors.New(errStr)
	}
	return outStr, nil
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
