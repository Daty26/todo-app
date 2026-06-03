package core_http_server

import (
	"fmt"
	core_http_middleware "github.com/Daty26/todo-app/internal/core/transport/http/middleware"
	"net/http"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("v1")
	ApiVersion2 = ApiVersion("v2")
	ApiVersion3 = ApiVersion("v3")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
	middleware []core_http_middleware.Middleware
}

func NewAPIVersionRouter(version ApiVersion, middleware ...core_http_middleware.Middleware) *APIVersionRouter {
	return &APIVersionRouter{ServeMux: http.NewServeMux(), apiVersion: version}
}

func (r *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		r.Handle(pattern, route.WithMiddleware())
	}
}
func (r *APIVersionRouter) WithMiddleWare() http.Handler {
	return core_http_middleware.ChainMiddleware(
		r, r.middleware...)
}
