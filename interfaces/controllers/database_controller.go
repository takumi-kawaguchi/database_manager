package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetDatabases returns all databases list
func GetDatabases(c echo.Context) (err error) {
	// db := infrastructure.DBConnect()
	// defer db.Close()
	// var databases []domain.Database
	// db.Find(&databases)
	return c.Render(http.StatusOK, "databases.tmpl", "databases")
}
