package definition

import "fmt"

// Route is a single route definition. It contains the necessary information to instantiate a route listener in the server package
type Route struct {
	Path        string                  `json:"path"`
	StatusCode  int                     `json:"statusCode"`
	Fixture     *map[string]interface{} `json:"fixture"`
	FixturePath string                  `json:"fixturePath"`
	Methods     []string                `json:"methods"`
	Headers     map[string]string       `json:"headers"`
}

func (r Route) String() string {
	return fmt.Sprintf("Route:\n\tPath: %v\n\tStatusCode: %v\n\tFixture: %v\n\tMethods: %v\n\tHeaders: %v", r.Path, r.StatusCode, r.Fixture, r.Methods, r.Headers)
}
