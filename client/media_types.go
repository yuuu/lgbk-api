// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "lgbk": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/yuuu/lgbk-api/design
// --out=$(GOPATH)/src/github.com/yuuu/lgbk-api
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// A date in lgbk (default view)
//
// Identifier: application/vnd.goa.example.date+json; view=default
type GoaExampleDate struct {
	// body of date
	Body string `form:"body" json:"body" yaml:"body" xml:"body"`
	// Unique bottle ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// title of date
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
}

// Validate validates the GoaExampleDate media type instance.
func (mt *GoaExampleDate) Validate() (err error) {

	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
	}
	return
}

// DecodeGoaExampleDate decodes the GoaExampleDate instance encoded in resp body.
func (c *Client) DecodeGoaExampleDate(resp *http.Response) (*GoaExampleDate, error) {
	var decoded GoaExampleDate
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
