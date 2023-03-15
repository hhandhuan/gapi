package consts

import "errors"

var (
	NoRouteErr      = errors.New("no route error")
	NoMethodErr     = errors.New("no method error")
	NoAuthFailedErr = errors.New("authorization failed")
)
