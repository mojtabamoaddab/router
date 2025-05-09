/*
Simple HTTP router built on Go 1.22's http.ServeMux, designed for easy route management
and middleware support.
*/
package router

import (
	"net/http"
)

// Router represents a router for handling HTTP requests.
type Router struct {
	root *Route
	RouterMethodsMixin
}

// New creates a new Router instance.
func New() *Router {
	return &Router{
		root: NewRoute(),
	}
}

// Use appends middlewares to the router.
func (r *Router) Use(middlewares ...Middleware) {
	r.root.chain.Append(middlewares...)
}

// Handle registers a handler for the given pattern.
func (r *Router) Handle(pattern string, handler http.Handler) {
	r.root.Handle(pattern, handler)
}

// HandleFunc registers a handler for the given pattern using a function.
func (r *Router) HandleFunc(pattern string, f func(http.ResponseWriter, *http.Request)) {
	r.root.Handle(pattern, http.HandlerFunc(f))
}

// Route registers a subroute for the given pattern.
func (r *Router) Route(pattern string, route *Route) {
	r.root.Route(pattern, route)
}

// RouteFunc registers a subroute for the given pattern using a function.
func (r *Router) RouteFunc(pattern string, f func(*Route)) {
	r.root.RouteFunc(pattern, f)
}

// ServeMux create an http.ServeMux instance and applies registered handlers.
func (r *Router) ServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	for _, h := range r.root.handlers {
		mux.Handle(h.pattern, h.handler)
	}

	return mux
}
