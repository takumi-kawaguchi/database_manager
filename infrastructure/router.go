package infrastructure

import (
	"text/template"

	"github.com/labstack/echo"
	"github.com/takumi-kawaguchi/database_manager/interfaces/controllers"
)

// Init initializes route config
func Init() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.gotmpl")),
	}
	e := echo.New()
	e.Renderer = t

	// router setting
	e.GET("/", func(c echo.Context) error { return controllers.GetDatabases(c) })

	e.Logger.Fatal(e.Start(":1323"))
}
