package hihttp

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tradjick/hiconfig"
)

func Serve(configPath string, configLoader hiconfig.ConfigLoader, routes []HiRouteType) {
	initUUID()
	e := echo.New()
	setGlobalMiddleware(e)
	mapRoutes(e, routes)
	e.Logger.Fatal(
		e.Start(
			fetchConfig(configPath, configLoader).AsString()))
}

func setGlobalMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//strip any injected request id and generate new
	e.Use(stripRequestID)
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: generateRequestID,
	}))
}
