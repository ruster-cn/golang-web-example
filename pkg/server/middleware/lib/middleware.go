package lib

import (
	"net/http"
	"path"
	"regexp"

	"github.com/gin-gonic/gin"
)

// Middleware receives a handlers and returns another handlers.
// The returned handlers can do some customized task according to
// the requirement
type Middleware func(http.Handler) http.Handler

// Chain make middleware together
func Chain(middlewares ...Middleware) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			h = middlewares[i](h)
		}

		return h
	}
}

// WithMiddlewares apply the middleware to the handlers.
// The middleware are executed in the order that they are applied
func WithMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	return Chain(middlewares...)(handler)
}

// New make a middleware from fn which type is func(w http.ResponseWriter, r *http.Request, next http.Handler)
func New(fn func(ctx *gin.Context), skippers ...Skipper) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, skipper := range skippers {
			if skipper(ctx.Request) {
				ctx.Next()
				return
			}
		}
		fn(ctx)
	}
}

// Skipper defines a function to skip middleware.
// Returning true skips processing the middleware.
type Skipper func(*http.Request) bool

// MethodAndPathSkipper returns skipper which
// will skip the middleware when r.Method equals the method and r.URL.Path matches the re
// when method is "*" it equals all http method
func MethodAndPathSkipper(method string, re *regexp.Regexp) func(r *http.Request) bool {
	return func(r *http.Request) bool {
		path := path.Clean(r.URL.EscapedPath())
		if (method == "*" || r.Method == method) && re.MatchString(path) {
			return true
		}

		return false
	}
}

// NegativeSkipper returns skipper which is negative of the input skipper
func NegativeSkipper(skipper Skipper) func(*http.Request) bool {
	return func(r *http.Request) bool {
		return !skipper(r)
	}
}
