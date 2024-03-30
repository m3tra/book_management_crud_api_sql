package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// import (
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

var (
	db *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("sqlite3", "fporto@Password123!@/book_mgmt?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
