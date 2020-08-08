package definition

import "fmt"

// Definition is a server definition. It contains all the information needed for the server package to create an HTTP server
type Definition struct {
	Port   int     `json:"port"`
	Routes []Route `json:"routes"`
}

func (d *Definition) String() string {
	return fmt.Sprintf("port: %v\nroutes: %v", d.Port, d.Routes)
}
