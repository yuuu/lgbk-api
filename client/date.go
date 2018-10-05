// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "lgbk": date Resource Client
//
// Command:
// $ goagen
// --design=github.com/yuuu/lgbk-api/design
// --out=$(GOPATH)/src/github.com/yuuu/lgbk-api
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// IndexDatePath computes a request path to the index action of date.
func IndexDatePath() string {

	return fmt.Sprintf("/dates/")
}

// Get all dates
func (c *Client) IndexDate(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewIndexDateRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewIndexDateRequest create the request corresponding to the index action endpoint of the date resource.
func (c *Client) NewIndexDateRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.BasicAuthSigner != nil {
		if err := c.BasicAuthSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ShowDatePath computes a request path to the show action of date.
func ShowDatePath(dateID int) string {
	param0 := strconv.Itoa(dateID)

	return fmt.Sprintf("/dates/%s", param0)
}

// Get date by id
func (c *Client) ShowDate(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowDateRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowDateRequest create the request corresponding to the show action endpoint of the date resource.
func (c *Client) NewShowDateRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.BasicAuthSigner != nil {
		if err := c.BasicAuthSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}