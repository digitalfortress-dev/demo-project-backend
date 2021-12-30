package common

import "github.com/labstack/echo/v4"

type (
	Route struct {
		Method  string
		Path    string
		Handler echo.HandlerFunc
	}

	Controller interface {
		Routes() []Route
	}
)
