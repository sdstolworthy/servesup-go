package definition

import "fmt"

type Route struct {
	Path       string   `json:"path"`
	StatusCode int32    `json:"statusCode"`
	Fixture    string   `json:"fixture"`
	Methods    []string `json:"methods"`
}

func (r Route) String() string {
	return fmt.Sprintf("Route:\n\tPath: %v\n\tStatusCode: %v\n\tFixture: %v\n\tMethods: %v\n", r.Path, r.StatusCode, r.Fixture, r.Methods)
}
