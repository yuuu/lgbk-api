// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "lgbk": Application Security
//
// Command:
// $ goagen
// --design=github.com/yuuu/lgbk-api/design
// --out=$(GOPATH)/src/github.com/yuuu/lgbk-api
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

type (
	// Private type used to store auth handler info in request context
	authMiddlewareKey string
)

// UseBasicAuthMiddleware mounts the basic_auth auth middleware onto the service.
func UseBasicAuthMiddleware(service *goa.Service, middleware goa.Middleware) {
	service.Context = context.WithValue(service.Context, authMiddlewareKey("basic_auth"), middleware)
}

// NewBasicAuthSecurity creates a basic_auth security definition.
func NewBasicAuthSecurity() *goa.BasicAuthSecurity {
	def := goa.BasicAuthSecurity{}
	return &def
}

// handleSecurity creates a handler that runs the auth middleware for the security scheme.
func handleSecurity(schemeName string, h goa.Handler, scopes ...string) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		scheme := ctx.Value(authMiddlewareKey(schemeName))
		am, ok := scheme.(goa.Middleware)
		if !ok {
			return goa.NoAuthMiddleware(schemeName)
		}
		ctx = goa.WithRequiredScopes(ctx, scopes)
		return am(h)(ctx, rw, req)
	}
}