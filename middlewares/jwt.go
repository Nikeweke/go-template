package middlewares

import (
	"github.com/labstack/echo/v4"
	"go-template/vars"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)


func JWT() echo.MiddlewareFunc {
	return echoMiddleware.JWTWithConfig(
		echoMiddleware.JWTConfig{
			SigningKey: []byte(vars.Config.JWT_SECRET),
		},
	)
}