package document

import (
	"encoding/json"
	"io/ioutil"
	"os"

	db "github.com/faruqfadhil/bibik/internal/database"
)

type document struct {
	path string
}

func NewDocument(path string) db.Database {
	return &document{
		path: path,
	}
}

func (d *document) Save(req *db.Model) error {
	payload := map[string]string{}
	if _, err := os.Stat(d.path); os.IsNotExist(err) {
		// create file if not exist.
		_, e := os.Create(d.path)
		if e != nil {
			return e
		}
	}

	// get existing data
	existingData, err := ioutil.ReadFile(d.path)
	if err != nil {
		return err
	}

	if len(existingData) > 0 {
		if err := json.Unmarshal(existingData, &payload); err != nil {
			return err
		}
	}

	// append with new data.
	payload[req.Key] = req.Value

	bytesData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// save again to file
	err = ioutil.WriteFile(d.path, bytesData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (d *document) FindByKey(key string) (*db.Model, error) {
	// get existing data
	existingData, err := ioutil.ReadFile(d.path)
	if err != nil {
		return nil, err
	}

	if len(existingData) == 0 {
		return nil, nil
	}

	existingDataInMap := map[string]string{}

	if err := json.Unmarshal(existingData, &existingDataInMap); err != nil {
		return nil, err
	}

	if val, ok := existingDataInMap[key]; ok {
		return &db.Model{
			Key:   key,
			Value: val,
		}, nil
	}

	return nil, nil
}
