package domain

import "github.com/jinzhu/gorm"

// Column structs infomation about column
type Column struct {
	gorm.Model
	PhysicalName     string
	LogicalName      string
	DataType         int
	Byte             int
	IsNotnull        bool
	IsPrimaryKey     bool
	IsUniqueIndexKey bool
	IsEncrypted      bool
	Index            int
	Description      string
	TableRefer       uint
}
