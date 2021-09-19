package database_test

import (
	"testing"

	bibikDB "github.com/faruqfadhil/bibik/internal/database"
	"github.com/faruqfadhil/bibik/internal/database/document"
	"github.com/google/go-cmp/cmp"
)

func initContainer(path string) bibikDB.Database {
	db := document.NewDocument(path)
	return db
}

func TestSave(t *testing.T) {
	db := initContainer("./document/data/bibik.data")

	tests := map[string]struct {
		req *bibikDB.Model
		err error
	}{
		"success": {
			req: &bibikDB.Model{
				Key:   "key 1",
				Value: "value 1",
			},
			err: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := db.Save(&bibikDB.Model{
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
	db := initContainer("./document/data/bibik.data")
	tests := map[string]struct {
		req string
		out *bibikDB.Model
		err error
	}{
		"success": {
			req: "key 1",
			out: &bibikDB.Model{
				Key:   "key 1",
				Value: "value 1",
			},
			err: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := db.FindByKey(test.req)
			if diff := cmp.Diff(out, test.out); diff != "" {
				t.Fatalf("output diff: expect: %v, got: %v", test.out, out)
			}
			if err != test.err {
				t.Fatalf("err diff: expect: %v, got: %v", test.err, err)
			}
		})
	}
}
