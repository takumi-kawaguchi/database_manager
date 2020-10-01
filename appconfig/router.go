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
	e.GET("/", func(c echo.Context) error { return controllers.GetDatabases(c) })
	e.GET("/database/new/", func(c echo.Context) error { return controllers.NewDatabase(c) })
	e.GET("/database/:database_id/tables/", func(c echo.Context) error { return controllers.GetTables(c) })
	e.GET("/database/:database_id/table/:table_id/", func(c echo.Context) error { return controllers.GetTable(c) })

	e.HTTPErrorHandler = infrastructure.JSONErrorHandler

	e.Start(":1324")
}
