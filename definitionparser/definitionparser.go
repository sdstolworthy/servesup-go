package definitionparser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/sdstolworthy/servesup/definition"
)

func resolveFilePath(basePath string, filePath string) (*string, error) {
	resolvedPath := path.Join(basePath, filePath)
	return &resolvedPath, nil
}

func parseRouteFixture(route definition.Route, definitionFilePath string) map[string]interface{} {
	if route.Fixture != nil {
		return *route.Fixture
	} else if route.FixturePath != "" {
		fixturePath, _ := resolveFilePath(definitionFilePath, route.FixturePath)
		jsonFile, _ := os.Open(*fixturePath)
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		fixture := map[string]interface{}{}
		json.Unmarshal(byteValue, &fixture)
		fmt.Println(*fixturePath)
		return fixture
	}
	return map[string]interface{}{}
}

func loadJSONFile(definitionPath string) *definition.Definition {
	jsonFile, err := os.Open(definitionPath)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	var definition definition.Definition
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &definition)
	for index, route := range definition.Routes {
		fixture := parseRouteFixture(route, path.Dir(definitionPath))
		definition.Routes[index].Fixture = &fixture
	}
	return &definition
}

type DoesNotExistError struct {
	s string
}

func (e *DoesNotExistError) Error() string {
	return e.s
}

func ParseDefinitionFromPath(path string) (*definition.Definition, error) {
	basePath, _ := os.Getwd()
	resolvedPath, _ := resolveFilePath(basePath, path)
	_, err := os.Stat(*resolvedPath)
	if os.IsNotExist(err) {
		return nil, &DoesNotExistError{s: fmt.Sprintf("%v does not exist", path)}
	}
	definition := loadJSONFile(*resolvedPath)
	return definition, nil
}
