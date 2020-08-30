package infrastructure

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DBConnect returns a DB struct for build connection to databases.
func DBConnect() *gorm.DB {
	con := "root:root@tcp(mysql-container:3306)/DBManagerDB?parseTime=true"
	db, err := gorm.Open("mysql", con)
	if err != nil {
		panic(err.Error())
	}
	return db
}
