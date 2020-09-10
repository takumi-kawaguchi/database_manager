package appcofig

import (
	"github.com/labstack/echo"
	"github.com/takumi-kawaguchi/database_manager/infrastructure"
	"github.com/takumi-kawaguchi/database_manager/interfaces/controllers"
)

// Init initializes route config
func Init() {
	e := echo.New()
	e.Renderer = infrastructure.NewTemplates()

	// router setting

	// GET / : databases list
	e.GET("/", func(c echo.Context) error { return controllers.GetDatabases(c) })
	// GET /database/new : new database form
	e.GET("/database/new/", func(c echo.Context) error { return controllers.NewDatabase(c) })

	e.Start(":1324")
}
