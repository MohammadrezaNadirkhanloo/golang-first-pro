package database

import "gorm.io/gorm"

func Connections() *gorm.DB {
	return DB
}