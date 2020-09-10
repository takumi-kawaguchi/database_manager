package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/takumi-kawaguchi/database_manager/domain"
	"github.com/takumi-kawaguchi/database_manager/infrastructure"
)

// GetDatabases returns all databases list and renders them
func GetDatabases(c echo.Context) (err error) {
	db := infrastructure.DBConnect()
	defer db.Close()
	var databases []domain.Database
	db.Find(&databases)
	return c.Render(http.StatusOK, "databases.tmpl", echo.Map{
		"databases": databases,
	})
}

// NewDatabase creates new database page
func NewDatabase(c echo.Context) (err error) {
	db := infrastructure.DBConnect()
	defer db.Close()
	return c.Render(http.StatusOK, "new_database.tmpl", nil)
}
