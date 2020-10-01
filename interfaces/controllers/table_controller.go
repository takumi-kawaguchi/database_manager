package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/takumi-kawaguchi/database_manager/domain"
	"github.com/takumi-kawaguchi/database_manager/infrastructure"
)

// GetTables returns all tables list and renders them
func GetTables(c echo.Context) error {
	db := infrastructure.DBConnect()
	defer db.Close()
	var tables []domain.Table
	db.Find(&tables, "database_id=?", c.Param("database_id"))
	return c.Render(http.StatusOK, "tables.tmpl", echo.Map{
		"tables": tables,
	})
}

// GetTable returns a table detail
func GetTable(c echo.Context) error {
	db := infrastructure.DBConnect()
	defer db.Close()
	var table domain.Table
	var columns []domain.Column
	db.Find("id=?", c.Param("table_id")).First(&table)
	db.Where("table_id=?", c.Param("table_id")).Find(&columns)
	fmt.Println("ここまで")
	test := c.Render(http.StatusOK, "table.tmpl", echo.Map{
		"table":   table,
		"columns": columns,
	})
	fmt.Println("reached here")
	fmt.Println(test)
	return test
}
