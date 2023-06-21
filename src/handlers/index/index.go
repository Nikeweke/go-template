package index 

import (
	"github.com/labstack/echo/v4"
	"go-template/services/utils"
)

func Index(c echo.Context) error {
	return utils.OkRequest(c, echo.Map{
		"Hello": "index",
	})
}

func Index2(c echo.Context) error {
	return utils.OkRequest(c, echo.Map{
		"Hello": "index2",
	})
}

func Index3(c echo.Context) error {
	return utils.OkRequest(c, echo.Map{
		"Hello": "index3",
	})
}