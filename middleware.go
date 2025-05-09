package router

import (
	"net/http"
)

// MiddlewareChain is a slice of Middleware.
type MiddlewareChain []Middleware

// Apply applies the middlewares in the chain to the given handler and returns a new http.Handler.
func (mc MiddlewareChain) Apply(h http.Handler) http.Handler {
	for i := len(mc) - 1; i >= 0; i-- {
		h = mc[i].Middleware(h)
	}
	return h
}

// Append adds middlewares to the chain.
func (mc *MiddlewareChain) Append(middlewares ...Middleware) {
	*mc = append(*mc, middlewares...)
}

// Copy creates a copy of the middleware chain.
func (mc MiddlewareChain) Copy() MiddlewareChain {
	newMC := make(MiddlewareChain, len(mc))
	copy(newMC, mc)
	return newMC
}

// Middleware interface defines a method that receives a http.Handler and returns another http.Handler.
type Middleware interface {
	Middleware(http.Handler) http.Handler
}

// MiddlewareFunc is a function type that allows ordinary functions to be used as Middleware.
type MiddlewareFunc func(http.Handler) http.Handler

// Middleware allows MiddlewareFunc to implement the middleware interface.
func (mf MiddlewareFunc) Middleware(h http.Handler) http.Handler {
	return mf(h)
}
