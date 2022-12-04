package hihttp

import "github.com/labstack/echo/v4"

type HiRouteType = func(e *echo.Echo)

func mapRoutes(e *echo.Echo, routes []HiRouteType) {
	for _, r := range routes {
		r(e)
	}
}
