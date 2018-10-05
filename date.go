package main

import (
	"github.com/goadesign/goa"
	"github.com/yuuu/lgbk-api/app"
)

// DateController implements the date resource.
type DateController struct {
	*goa.Controller
}

// NewDateController creates a date controller.
func NewDateController(service *goa.Service) *DateController {
	return &DateController{Controller: service.NewController("DateController")}
}

// Index runs the index action.
func (c *DateController) Index(ctx *app.IndexDateContext) error {
	// DateController_Index: start_implement

	// Put your logic here

	res := &app.GoaExampleDate{}
	return ctx.OK(res)
	// DateController_Index: end_implement
}

// Show runs the show action.
func (c *DateController) Show(ctx *app.ShowDateContext) error {
	// DateController_Show: start_implement

	// Put your logic here

	res := &app.GoaExampleDate{}
	return ctx.OK(res)
	// DateController_Show: end_implement
}
