package main

import (
	"github.com/labstack/echo/v4"

	"github.com/go-playground/validator"
	"html/template"
	"io"
)

// Example of template, you can remove it 
type Template struct {
	templates *template.Template
}
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
// var TemplateRenderer = &Template{
// 	templates: template.Must(template.ParseGlob("./public/views/*.html")),
// }

// Example of validator, you can remove it 
type CustomValidator struct {
	validator *validator.Validate
}
func (cv *CustomValidator) Validate(i interface{}) error {
	// switch v := i.(type) {
	// 	case *models.TaxDeclarationCreateRequest: // you must have this model with "CheckRequestBody" implemented
	// 		return v.CheckRequestBody()
	// 	default:
	// 		return nil
	// }
  return nil
}
var CustomValidatorInstance = &CustomValidator{validator: validator.New()}