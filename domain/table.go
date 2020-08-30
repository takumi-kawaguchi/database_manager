package domain

import "github.com/jinzhu/gorm"

// Table structs infomation about table
type Table struct {
	gorm.Model
	Name              string
	Description       string
	DatabaseRefer uint
	Columns           []Column `gorm:"foreignkey:TableRefer`
}
