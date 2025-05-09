package router

import (
	"net/http"
)

type handler interface {
	Handle(string, http.Handler)
}

// RouterMethodsMixin provides high-level API methods for HTTP methods like GET, POST, etc.
type RouterMethodsMixin struct {
	handler handler
}

// HandleFunc registers a handler function for the given pattern.
func (m *RouterMethodsMixin) HandleFunc(pattern string, f http.HandlerFunc) {
	m.handler.Handle(pattern, f)
}

// Method registers a handler for a specific HTTP method and pattern.
func (m *RouterMethodsMixin) Method(method string, pattern string, handler http.Handler) {
	m.handler.Handle(method+" "+pattern, handler)
}

// MethodFunc registers a handler function for a specific HTTP method and pattern.
func (m *RouterMethodsMixin) MethodFunc(method string, pattern string, f http.HandlerFunc) {
	m.Method(method, pattern, f)
}

// Get registers a handler for the GET method and pattern.
func (m *RouterMethodsMixin) Get(pattern string, handler http.Handler) {
	m.Method(http.MethodGet, pattern, handler)
}

// GetFunc registers a handler function for the GET method and pattern.
func (m *RouterMethodsMixin) GetFunc(pattern string, f http.HandlerFunc) {
	m.Get(pattern, f)
}

// Post registers a handler for the POST method and pattern.
func (m *RouterMethodsMixin) Post(pattern string, handler http.Handler) {
	m.Method(http.MethodPost, pattern, handler)
}

// PostFunc registers a handler function for the POST method and pattern.
func (m *RouterMethodsMixin) PostFunc(pattern string, f http.HandlerFunc) {
	m.Post(pattern, f)
}

// Patch registers a handler for the PATCH method and pattern.
func (m *RouterMethodsMixin) Patch(pattern string, handler http.Handler) {
	m.Method(http.MethodPatch, pattern, handler)
}

// PatchFunc registers a handler function for the PATCH method and pattern.
func (m *RouterMethodsMixin) PatchFunc(pattern string, f http.HandlerFunc) {
	m.Patch(pattern, f)
}

// Put registers a handler for the PUT method and pattern.
func (m *RouterMethodsMixin) Put(pattern string, handler http.Handler) {
	m.Method(http.MethodPut, pattern, handler)
}

// PutFunc registers a handler function for the PUT method and pattern.
func (m *RouterMethodsMixin) PutFunc(pattern string, f http.HandlerFunc) {
	m.Put(pattern, f)
}

// Delete registers a handler for the DELETE method and pattern.
func (m *RouterMethodsMixin) Delete(pattern string, handler http.Handler) {
	m.Method(http.MethodDelete, pattern, handler)
}

// DeleteFunc registers a handler function for the DELETE method and pattern.
func (m *RouterMethodsMixin) DeleteFunc(pattern string, f http.HandlerFunc) {
	m.Delete(pattern, f)
}

// Head registers a handler for the HEAD method and pattern.
func (m *RouterMethodsMixin) Head(pattern string, handler http.Handler) {
	m.Method(http.MethodHead, pattern, handler)
}

// HeadFunc registers a handler function for the HEAD method and pattern.
func (m *RouterMethodsMixin) HeadFunc(pattern string, f http.HandlerFunc) {
	m.Head(pattern, f)
}

// Options registers a handler for the OPTIONS method and pattern.
func (m *RouterMethodsMixin) Options(pattern string, handler http.Handler) {
	m.Method(http.MethodOptions, pattern, handler)
}

// OptionsFunc registers a handler function for the OPTIONS method and pattern.
func (m *RouterMethodsMixin) OptionsFunc(pattern string, f http.HandlerFunc) {
	m.Options(pattern, f)
}
