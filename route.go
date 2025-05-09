package router

import (
	"net/http"
)

type handlerEntry struct {
	pattern string
	handler http.Handler
}

// Route represents a route with middleware and handlers.
type Route struct {
	chain    MiddlewareChain
	handlers []handlerEntry
	RouterMethodsMixin
}

// NewRoute creates a new Route instance.
func NewRoute() *Route {
	return &Route{
		chain:    make(MiddlewareChain, 0),
		handlers: make([]handlerEntry, 0),
	}
}

// Use appends middlewares to the route.
func (r *Route) Use(middlewares ...Middleware) {
	r.chain.Append(middlewares...)
}

// Handle registers a handler for the given pattern.
func (r *Route) Handle(pattern string, handler http.Handler) {
	r.handlers = append(r.handlers, handlerEntry{
		pattern: pattern,
		handler: r.chain.Apply(handler),
	})
}

// HandleFunc registers a handler for the given pattern using a function.
func (r *Route) HandleFunc(pattern string, f func(http.ResponseWriter, *http.Request)) {
	r.Handle(pattern, http.HandlerFunc(f))
}

// Route registers a subroute for the given pattern.
func (r *Route) Route(pattern string, route *Route) {
	if err := validateRoutePattern(pattern); err != nil {
		panic(err)
	}
	for _, h := range route.handlers {
		r.Handle(pattern+h.pattern, h.handler)
	}
}

// RouteFunc registers a subroute for the given pattern using a function.
func (r *Route) RouteFunc(pattern string, f func(*Route)) {
	route := NewRoute()
	f(route)
	r.Route(pattern, route)
}
