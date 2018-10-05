// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "lgbk": Application Contexts
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
	"strconv"
)

// IndexDateContext provides the date index action context.
type IndexDateContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewIndexDateContext parses the incoming request URL and body, performs validations and creates the
// context used by the date controller index action.
func NewIndexDateContext(ctx context.Context, r *http.Request, service *goa.Service) (*IndexDateContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := IndexDateContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *IndexDateContext) OK(r *GoaExampleDate) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.date+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *IndexDateContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowDateContext provides the date show action context.
type ShowDateContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	DateID int
}

// NewShowDateContext parses the incoming request URL and body, performs validations and creates the
// context used by the date controller show action.
func NewShowDateContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowDateContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowDateContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramDateID := req.Params["dateID"]
	if len(paramDateID) > 0 {
		rawDateID := paramDateID[0]
		if dateID, err2 := strconv.Atoi(rawDateID); err2 == nil {
			rctx.DateID = dateID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("dateID", rawDateID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowDateContext) OK(r *GoaExampleDate) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.date+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowDateContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
