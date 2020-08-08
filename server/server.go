package server

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sdstolworthy/servesup/definition"
)

func createResponseHandler(handlerMethod func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes, route definition.Route) gin.IRoutes {
	return handlerMethod(route.Path, func(c *gin.Context) {
		var returnJson *map[string]interface{}
		if route.Fixture != nil {
			returnJson = route.Fixture
		} else {
			returnJson = &map[string]interface{}{}
		}
		c.JSON(route.StatusCode, returnJson)
	})
}

var methodMap = map[string]func(r *gin.Engine, route definition.Route){
	"get": func(r *gin.Engine, route definition.Route) {
		createResponseHandler(r.GET, route)
	},
	"head": func(r *gin.Engine, route definition.Route) {
		createResponseHandler(r.HEAD, route)
	},
	"delete": func(r *gin.Engine, route definition.Route) {
		createResponseHandler(r.DELETE, route)
	},
	"post": func(r *gin.Engine, route definition.Route) {
		createResponseHandler(r.POST, route)
	},
	"options": func(r *gin.Engine, route definition.Route) {
		createResponseHandler(r.OPTIONS, route)
	},
	"patch": func(r *gin.Engine, route definition.Route) {
		createResponseHandler(r.PATCH, route)
	},
	"put": func(r *gin.Engine, route definition.Route) {
		createResponseHandler(r.PUT, route)
	},
}

// RunServer accepts a server definition and intantiates an Http server that matches the definition spec
func RunServer(definition *definition.Definition) {
	router := gin.Default()

	for _, route := range definition.Routes {
		if route.Path == "" || route.Path == "*" || route.Path == "/*" {
			router.NoRoute(func(c *gin.Context) {
				c.JSON(route.StatusCode, gin.H{"no": "route"})
			})
			continue
		}
		if len(route.Methods) == 0 {
			router.Any(route.Path)
			continue
		}
		for _, method := range route.Methods {
			if val, ok := methodMap[strings.ToLower(method)]; ok {
				val(router, route)
			}
		}
	}
	router.Run(fmt.Sprintf(":%v", definition.Port))
}
