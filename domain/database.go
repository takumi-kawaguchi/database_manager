package domain

import "github.com/jinzhu/gorm"

// Database structs infomation about database
type Database struct {
	gorm.Model
	Name         string
	Description  string
	ControllType int
	Tables       []Table `gorm:"foreignkey:DatabaseRefer`
}
