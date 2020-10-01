package infrastructure

import (
	"errors"
	"io"
	"text/template"

	"github.com/labstack/echo"
)

// Template has template info
type Template struct {
	templates *template.Template
}

// TemplateRepository has template list
type TemplateRepository struct {
	templates map[string]*template.Template
}

// Render renders a view
// func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
// 	return t.templates.ExecuteTemplate(w, name, data)
// }

// Render renders a view
func (t *TemplateRepository) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)

}

// NewTemplates returns templates
func NewTemplates() *TemplateRepository {
	templates := make(map[string]*template.Template)
	// TODO: ここの定義長ったらしくなりそうだから関数化したい
	templates["databases.tmpl"] = template.Must(template.ParseFiles("views/databases.tmpl", "views/base.tmpl"))
	templates["tables.tmpl"] = template.Must(template.ParseFiles("views/tables.tmpl", "views/base.tmpl"))
	templates["table.tmpl"] = template.Must(template.ParseFiles("views/table.tmpl", "views/base.tmpl"))
	// templates["new_database.tmpl"] = template.Must(template.ParseFiles("views/new_database.tmpl", "views/base.tmpl"))
	return &TemplateRepository{
		templates: templates,
	}
}
