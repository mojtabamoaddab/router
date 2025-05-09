package router

import (
	"errors"
	"strings"
)

var (
	errRouteMethod = errors.New("Route(...) pattern cannot include Method")
	errRouteSlash  = errors.New("Route(...) pattern must be ended with '/'")
)

func validateRoutePattern(pattern string) error {
	if len(pattern) < 1 || pattern[len(pattern)-1] != '/' {
		return errRouteSlash
	}
	if strings.Contains(pattern, " ") {
		return errRouteMethod
	}
	return nil
}
