package main

import (
	"fmt"

	"github.com/sdstolworthy/servesup/cmdparse"
	"github.com/sdstolworthy/servesup/definitionparser"
)

func main() {
	config := cmdparse.ParseCmdOptions()

	definition, err := definitionparser.ParseDefinitionFromPath(config.Definition)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(definition)
}
