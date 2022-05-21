package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dub *gorm.DB
)

func Connect() {
	dsn := "root@tcp(127.0.0.1:3306)"

	mysql.Open(dsn)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nill {
		panic(err)
	}
	db = d
}
func GgtDB() *gorm.DB {
	return db

}
