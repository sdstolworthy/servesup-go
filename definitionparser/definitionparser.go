package definitionparser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sdstolworthy/servesup/definition"
)

func loadJSONFile(path string) *definition.Definition {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	var definition definition.Definition
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &definition)
	return &definition
}

type DoesNotExistError struct {
	s string
}

func (e *DoesNotExistError) Error() string {
	return e.s
}

func ParseDefinitionFromPath(path string) (*definition.Definition, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, &DoesNotExistError{s: fmt.Sprintf("%v does not exist", path)}
	}
	definition := loadJSONFile(path)
	return definition, nil
}
