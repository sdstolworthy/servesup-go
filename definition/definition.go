package definition

import "fmt"

type Definition struct {
	Port   int32   `json:"port"`
	Routes []Route `json:"routes"`
}

func (d *Definition) String() string {
	return fmt.Sprintf("port: %v\nroutes: %v", d.Port, d.Routes)
}
