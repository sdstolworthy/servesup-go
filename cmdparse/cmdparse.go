package cmdparse

import (
	"flag"
	"fmt"
)

type ParsedOptions struct {
	Port       int
	Definition string
}

func (p ParsedOptions) String() string {
	return fmt.Sprintf("port: %v, definition: %v", p.Port, p.Definition)
}

func ParseCmdOptions() ParsedOptions {
	portPtr := flag.Int("port", 3000, "The port on which the server should listen")
	serverFilePtr := flag.String("definition", "", "The location of the server definition")
	flag.Parse()
	parsedOptions := ParsedOptions{Port: *portPtr, Definition: *serverFilePtr}
	return parsedOptions
}
