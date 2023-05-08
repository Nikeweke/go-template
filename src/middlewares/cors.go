package middlewares

import (
	"net/http"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Cors() echo.MiddlewareFunc {
	return echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
				http.MethodGet, 
				http.MethodHead, 
				http.MethodPut, 
				http.MethodPatch, 
				http.MethodPost, 
				http.MethodDelete,
		},
	})
}