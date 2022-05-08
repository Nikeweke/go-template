package handlers 

import (
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return OkRequest(c, map[string]string{
		"Hello": "index",
	})
}

func Index2(c echo.Context) error {
	return OkRequest(c, map[string]string{
		"Hello": "index2",
	})
}

func Index3(c echo.Context) error {
	return OkRequest(c, map[string]string{
		"Hello": "index3",
	})
}