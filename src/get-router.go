package main 

import (
	// "net/http"

	"go-template/handlers/index"

	"go-template/middlewares"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func GetRouter() *echo.Echo {
	r := echo.New()
	r.HideBanner = true 
	r.HidePort   = true

	// Global middlewares
	r.Use(echoMiddleware.Recover())    
	r.Use(middlewares.Cors())            
	r.Use(middlewares.Gzip()) 

	// Template example, you can remove it
	// additional-features.go
	// r.Renderer = TemplateRenderer

	// Validator example, you can remove it
	// additional-features.go
	// r.Validator = CustomValidatorInstance

	// static files example 
	r.Static("/assets", "public")

	// single route 
	r.GET("/", index.Index)

	// grouped route 
	routeGroup := r.Group("/rg") 
	routeGroup.GET("", index.Index2) // full route: /rg

	// grouped route protected
	routeGroup2 := r.Group("/rg2", middlewares.JWT()) 
	routeGroup2.GET("/protected", index.Index3) // full route: /rg2/protected
 
	return r
}
