// Code generated by go-swagger; DO NOT EDIT.

package attempt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// UserAttemptGetURL generates an URL for the user attempt get operation
type UserAttemptGetURL struct {
	StatusID *uint64
	TestID   *uint64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *UserAttemptGetURL) WithBasePath(bp string) *UserAttemptGetURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *UserAttemptGetURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *UserAttemptGetURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/user/attempt"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var statusIDQ string
	if o.StatusID != nil {
		statusIDQ = swag.FormatUint64(*o.StatusID)
	}
	if statusIDQ != "" {
		qs.Set("status_id", statusIDQ)
	}

	var testIDQ string
	if o.TestID != nil {
		testIDQ = swag.FormatUint64(*o.TestID)
	}
	if testIDQ != "" {
		qs.Set("test_id", testIDQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *UserAttemptGetURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *UserAttemptGetURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *UserAttemptGetURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on UserAttemptGetURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on UserAttemptGetURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *UserAttemptGetURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
