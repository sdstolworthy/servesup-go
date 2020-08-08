package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/sdstolworthy/servesup/definition"
)

func createResponseHandler(handlerMethod func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes, route definition.Route) gin.IRoutes {
	return handlerMethod(route.Path, func(c *gin.Context) {
		c.JSON(route.StatusCode, gin.H{"ping": "pong"})
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
}

func RunServer(definition *definition.Definition) {
	r := gin.Default()
	for _, route := range definition.Routes {
		if methodMap == nil {
			r.Any(route.Path)
			continue
		}
		for _, method := range route.Methods {
			if val, ok := methodMap[method]; ok {
				val(r, route)
			}
		}
	}
	r.Run(fmt.Sprintf(":%v", definition.Port))
}
