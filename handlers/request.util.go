package handlers 

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func BadRequest(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusBadRequest, data)
}

func OkRequest(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}