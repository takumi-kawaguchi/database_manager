package infrastructure

import (
	"io"
	"text/template"

	"github.com/labstack/echo"
)

// Template has template info
type Template struct {
	templates *template.Template
}

// Render renders a view
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// NewTemplate returns template
func NewTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("views/*.tmpl")),
	}
}
