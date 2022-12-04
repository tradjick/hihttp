package hihttp

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func generateRequestID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return u.String()
}

func initUUID() {
	uuid.EnableRandPool()
}

func stripRequestID(n echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Del(echo.HeaderXRequestID)
		return n(c)
	}
}
