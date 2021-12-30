package routes

import (
	"main/common"
	"main/patient"

	"github.com/labstack/echo/v4"
)

func DefineApiRoute(e *echo.Echo) {

	controllers := []common.Controller{patient.PatientController{}}
	var routes []common.Route
	for _, controller := range controllers {
		routes = append(routes, controller.Routes()...)
	}
	api := e.Group("/api")
	admin_api := e.Group("/api/admin")
	e.Static("/api/upload-photo", "patient")
	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			{
				api.POST(route.Path, route.Handler)
				break
			}
		case echo.GET:
			{
				admin_api.GET(route.Path, route.Handler, route.Middleware...)
				break
			}
		}
	}
}
