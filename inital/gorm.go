package inital

import (
	"fmt"
	"github.com/akazwz/hot-search-notify/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	m := CFG

	if m.Username == "" {
		return nil
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.DBName,
	)

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return nil
	} else {
		//sqlDB, _ := db.DB()
		//sqlDB.SetMaxIdleConns()
		//sqlDB.SetMaxIdleConns()
		return db
	}
}

func CreateTables(db *gorm.DB) {
	err := db.AutoMigrate(model.Notify{})
	if err != nil {
		log.Println("create table error")
	}
}
