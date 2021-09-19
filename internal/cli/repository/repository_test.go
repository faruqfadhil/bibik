package repository_test

import (
	"errors"
	"testing"

	"github.com/faruqfadhil/bibik/internal/cli/repository"
	"github.com/faruqfadhil/bibik/internal/repository/cli-repository/document"

	errLib "github.com/faruqfadhil/bibik/pkg/error"
	"github.com/google/go-cmp/cmp"
)

func initDocumentContainer() repository.CLIRepository {
	db := document.NewDocument("../../repository/cli-repository/document/data/bibik.data")
	return db
}

func TestUpsert(t *testing.T) {
	db := initDocumentContainer()

	tests := map[string]struct {
		req *repository.CLIModel
		err error
	}{
		"success": {
			req: &repository.CLIModel{
				Key:   "key 1",
				Value: "value 1",
			},
			err: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := db.Upsert(&repository.CLIModel{
				Key:   test.req.Key,
				Value: test.req.Value,
			})
			if err != test.err {
				t.Fatalf("err diff: expect: %v, got: %v", test.err, err)
			}
		})
	}
}

func TestFindByKey(t *testing.T) {
	db := initDocumentContainer()
	tests := map[string]struct {
		req string
		out *repository.CLIModel
		err error
	}{
		"success": {
			req: "key 1",
			out: &repository.CLIModel{
				Key:   "key 1",
				Value: "value 1",
			},
			err: nil,
		},
		"key not found": {
			req: "key 9999",
			out: nil,
			err: errLib.ErrKeyNotFound,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := db.FindByKey(test.req)
			if diff := cmp.Diff(out, test.out); diff != "" {
				t.Fatalf("output diff: expect: %v, got: %v", test.out, out)
			}
			if !errors.Is(test.err, err) {
				t.Fatalf("err diff: expect: %v, got: %v", test.err, err)
			}
		})
	}
}

func TestDeleteByKey(t *testing.T) {
	db := initDocumentContainer()
	tests := map[string]struct {
		req string
		err error
	}{
		"success": {
			req: "key 1",
			err: nil,
		},
		"key not found": {
			req: "key 9999",
			err: errLib.ErrKeyNotFound,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := db.DeleteByKey(test.req)
			if !errors.Is(test.err, err) {
				t.Fatalf("err diff: expect: %v, got: %v", test.err, err)
			}
		})
	}
}
