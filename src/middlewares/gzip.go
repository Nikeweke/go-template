package middlewares

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Gzip() echo.MiddlewareFunc {
	return echoMiddleware.GzipWithConfig(
		echoMiddleware.GzipConfig{
			Level: 5,
		},
	)
}