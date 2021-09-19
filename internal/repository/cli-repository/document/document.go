package document

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/faruqfadhil/bibik/internal/cli/repository"
	errLib "github.com/faruqfadhil/bibik/pkg/error"
)

type document struct {
	homedir  string
	path     string
	fullPath string
}

func NewDocument(homedir, path, fullPath string) repository.CLIRepository {
	return &document{
		homedir:  homedir,
		path:     path,
		fullPath: fullPath,
	}
}

func (d *document) validatePath() error {
	// check path to file
	if _, err := os.Stat(fmt.Sprintf("%s/%s", d.homedir, d.path)); os.IsNotExist(err) {
		// create dir if not exist
		err := os.Mkdir(fmt.Sprintf("%s/%s", d.homedir, d.path), 0755)
		if err != nil {
			return errLib.ErrCreateDir
		}
	}

	// check file
	if _, err := os.Stat(d.fullPath); os.IsNotExist(err) {
		// create file if not exist.
		_, e := os.Create(d.fullPath)
		if e != nil {
			return errLib.ErrCreateFile
		}
	}
	return nil
}

func (d *document) Upsert(req *repository.CLIModel) error {
	if err := d.validatePath(); err != nil {
		return err
	}

	// get existing data
	existingData, err := ioutil.ReadFile(d.fullPath)
	if err != nil {
		return errLib.ErrReadFile
	}

	payload := map[string]string{}
	if len(existingData) > 0 {
		if err := json.Unmarshal(existingData, &payload); err != nil {
			return errLib.ErrUnmarshal
		}
	}

	// append with new data.
	payload[req.Key] = req.Value
	return d.insert(payload)
}

func (d *document) FindByKey(key string) (*repository.CLIModel, error) {
	if err := d.validatePath(); err != nil {
		return nil, err
	}

	existingDataInMap, err := d.fetchAllExistingData()
	if err != nil {
		return nil, err
	}

	if val, ok := existingDataInMap[key]; ok {
		return &repository.CLIModel{
			Key:   key,
			Value: val,
		}, nil
	}

	return nil, errLib.ErrKeyNotFound

}

func (d *document) DeleteByKey(key string) error {
	if err := d.validatePath(); err != nil {
		return err
	}

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

func (d *document) SearchByKey(key string) ([]*repository.CLIModel, error) {
	if err := d.validatePath(); err != nil {
		return nil, err
	}

	existingDataInMap, err := d.fetchAllExistingData()
	if err != nil {
		return nil, err
	}

	out := []*repository.CLIModel{}
	for k, value := range existingDataInMap {
		if strings.Contains(strings.ToLower(k), strings.ToLower(key)) {
			out = append(out, &repository.CLIModel{
				Key:   k,
				Value: value,
			})
		}
	}

	if len(out) < 1 {
		return nil, errLib.ErrKeyNotFound
	}

	return out, nil
}

func (d *document) insert(payload map[string]string) error {
	bytesData, err := json.Marshal(payload)
	if err != nil {
		return errLib.ErrMarshal
	}

	// save again to file
	err = ioutil.WriteFile(d.fullPath, bytesData, 0644)
	if err != nil {
		return errLib.ErrWriteFile
	}

	return nil
}

func (d *document) fetchAllExistingData() (map[string]string, error) {
	existingData, err := ioutil.ReadFile(d.fullPath)
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
