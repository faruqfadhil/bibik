package document_test

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/faruqfadhil/bibik/internal/cli/repository"
	"github.com/faruqfadhil/bibik/internal/repository/cli-repository/document"

	errLib "github.com/faruqfadhil/bibik/pkg/error"
	"github.com/stretchr/testify/assert"
)

func initDocumentContainer() repository.CLIRepository {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf(err.Error())
	}
	path := ".bibik"
	db := document.NewDocument(dirname, path, fmt.Sprintf("%s/%s/%s", dirname, path, "bibik.data"))
	return db
}

func TestUpsert(t *testing.T) {
	db := initDocumentContainer()

	tests := map[string]struct {
		req *repository.CLIModel
		err error
	}{
		"success insert data 1": {
			req: &repository.CLIModel{
				Key:   "key 1",
				Value: "value 1",
			},
			err: nil,
		},
		"success insert data 2": {
			req: &repository.CLIModel{
				Key:   "key 2",
				Value: "value 2",
			},
			err: nil,
		},
		"success insert data 3": {
			req: &repository.CLIModel{
				Key:   "number 1",
				Value: "val 1",
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
			assert.Equal(t, test.out, out, "output different")
			if !errors.Is(test.err, err) {
				t.Fatalf("err diff: expect: %v, got: %v", test.err, err)
			}
		})
	}
}

func TestSearchByKey(t *testing.T) {
	db := initDocumentContainer()
	tests := map[string]struct {
		req string
		out []*repository.CLIModel
		err error
	}{
		"success": {
			req: "key",
			out: []*repository.CLIModel{
				{
					Key:   "key 2",
					Value: "value 2",
				},
				{
					Key:   "key 1",
					Value: "value 1",
				},
			},
			err: nil,
		},
		"key not found": {
			req: "9999",
			out: nil,
			err: errLib.ErrKeyNotFound,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := db.SearchByKey(test.req)
			assert.ElementsMatch(t, out, test.out, "output different")
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
			req: "number 1",
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
