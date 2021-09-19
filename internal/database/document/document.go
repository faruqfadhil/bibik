package document

import (
	"encoding/json"
	"io/ioutil"
	"os"

	db "github.com/faruqfadhil/bibik/internal/database"
	errLib "github.com/faruqfadhil/bibik/pkg/error"
)

type document struct {
	path string
}

func NewDocument(path string) db.Database {
	return &document{
		path: path,
	}
}

func (d *document) Upsert(req *db.Model) error {
	payload := map[string]string{}
	if _, err := os.Stat(d.path); os.IsNotExist(err) {
		// create file if not exist.
		_, e := os.Create(d.path)
		if e != nil {
			return errLib.ErrCreateFile
		}
	}

	// get existing data
	existingData, err := ioutil.ReadFile(d.path)
	if err != nil {
		return errLib.ErrReadFile
	}

	if len(existingData) > 0 {
		if err := json.Unmarshal(existingData, &payload); err != nil {
			return errLib.ErrUnmarshal
		}
	}

	// append with new data.
	payload[req.Key] = req.Value
	return d.insert(payload)
}

func (d *document) FindByKey(key string) (*db.Model, error) {
	existingDataInMap, err := d.fetchAllExistingData()
	if err != nil {
		return nil, err
	}

	if val, ok := existingDataInMap[key]; ok {
		return &db.Model{
			Key:   key,
			Value: val,
		}, nil
	}

	return nil, errLib.ErrKeyNotFound

}

func (d *document) DeleteByKey(key string) error {
	existingDataInMap, err := d.fetchAllExistingData()
	if err != nil {
		return err
	}

	if _, ok := existingDataInMap[key]; !ok {
		return errLib.ErrKeyNotFound
	}
	delete(existingDataInMap, key)

	return d.insert(existingDataInMap)
}

func (d *document) insert(payload map[string]string) error {
	bytesData, err := json.Marshal(payload)
	if err != nil {
		return errLib.ErrMarshal
	}

	// save again to file
	err = ioutil.WriteFile(d.path, bytesData, 0644)
	if err != nil {
		return errLib.ErrWriteFile
	}

	return nil
}

func (d *document) fetchAllExistingData() (map[string]string, error) {
	existingData, err := ioutil.ReadFile(d.path)
	if err != nil {
		return nil, errLib.ErrReadFile
	}
	if len(existingData) == 0 {
		return nil, errLib.ErrDataEmpty
	}
	existingDataInMap := map[string]string{}

	if err := json.Unmarshal(existingData, &existingDataInMap); err != nil {
		return nil, errLib.ErrUnmarshal
	}
	return existingDataInMap, nil
}
